package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/haziqam/course-scheduler-backend/packages/algorithm"
	"github.com/haziqam/course-scheduler-backend/packages/database"
	"github.com/haziqam/course-scheduler-backend/packages/models"
)

func GetAllMatkul(c *fiber.Ctx) error {
	query := `
		SELECT * 
		FROM matkul
	`
	rows, err := database.DbInstance.Query(query)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": "Query error"})
	}

	defer rows.Close()

	var matkulArr []models.Matkul
	
	for rows.Next() {
		matkul := new(models.Matkul)
		err = matkul.ScanRow(rows);
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{"error": "Error scanning rows"})
		}
		matkulArr = append(matkulArr, *matkul)
	}

	err = rows.Err()
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": "Error iterating rows"})
	}

	return c.JSON(matkulArr)
}

func addMatkul(c *fiber.Ctx, newMatkul []models.Matkul) error {
	query := `
		INSERT INTO matkul(nama_matkul, sks, nama_jurusan, min_semester, prediksi) 
		VALUES ($1, $2, $3, $4, $5)
	`

	for _, matkul := range newMatkul {
		_, err := database.DbInstance.Exec(query, matkul.NamaMatkul, matkul.SKS, matkul.NamaJurusan, 
			matkul.MinSemester, matkul.PrediksiIndeks)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{"error": "Query failed"})
		}
	}
	
	c.Status(fiber.StatusCreated)
	return c.JSON(fiber.Map{"message": "Matkul added successfully"})
}

func AddMatkul(c *fiber.Ctx) error {
	var newMatkul []models.Matkul
	err := c.BodyParser(&newMatkul)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": "Error parsing request body"})
	}

	return addMatkul(c, newMatkul)
}

func AddMatkulFromFile(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		log.Println("Error parsing form:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error parsing form"})
	}

	files := form.File["Matkul[]"]
	if len(files) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No file uploaded"})
	}

	file, err := files[0].Open()
	if err != nil {
		log.Println("Error opening file:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error opening file"})
	}
	defer file.Close()

	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Error reading file:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error reading file"})
	}

	var newMatkul []models.Matkul
	err = json.Unmarshal(fileContent, &newMatkul)
	if err != nil {
		log.Println("Error unmarshaling file content:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error unmarshaling file content"})
	}

	return addMatkul(c, newMatkul)
}

func RemoveMatkul(c *fiber.Ctx) error {
		namaMatkul := c.Query("matkul")
		namaJurusan := c.Query("jurusan")
		query := `
			DELETE FROM matkul
			WHERE nama_matkul = ($1)
			AND nama_jurusan = ($2)
		`
	
		_, err := database.DbInstance.Exec(query, namaMatkul, namaJurusan)
	
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{"error": "Failed to delete matkul"})
		}
	
		return c.SendStatus(fiber.StatusNoContent)
}

func filterMatkul(namaFakultas string, currentSemester int) ([]models.Matkul, error) {
	query := `
	SELECT nama_matkul, sks, nama_jurusan, min_semester, prediksi
	FROM matkul NATURAL JOIN jurusan
	WHERE nama_fakultas = $1
	AND min_semester <= $2;
	`

	rows, err := database.DbInstance.Query(query, namaFakultas, currentSemester)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var filteredMatkul []models.Matkul

	for rows.Next() {
		matkul := new(models.Matkul)
		err = matkul.ScanRow(rows);
		if err != nil {
			return nil, err
		}
		filteredMatkul = append(filteredMatkul, *matkul)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return filteredMatkul, nil
}

func FindMatkul(c *fiber.Ctx) error {
	namaFakultas := c.Query("fakultas")
	currentSemester, err := strconv.Atoi(c.Query("semester"))

	if namaFakultas == "" || err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	filteredMatkul, err := filterMatkul(namaFakultas, currentSemester)

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(filteredMatkul)
}

func FindBestOptions(c *fiber.Ctx) error {
	namaFakultas := c.Query("fakultas")
	currentSemester, err := strconv.Atoi(c.Query("semester"))
	minSKS, err := strconv.Atoi(c.Query("minSKS"))
	maxSKS, err := strconv.Atoi(c.Query("maxSKS"))

	if namaFakultas == "" || err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	filteredMatkul, err := filterMatkul(namaFakultas, currentSemester)

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	if len(filteredMatkul) == 0 {
		c.Status(fiber.StatusNotFound)
		errorMessage := fmt.Sprintf("No matkul with nama fakultas = %s and minimum semester <= %d found", namaFakultas, currentSemester)
		return c.JSON(fiber.Map{"error": errorMessage})
	}

	if models.CountTotalSKS(filteredMatkul) < minSKS {
		c.Status(fiber.StatusNotFound)
		errorMessage := "Total SKS of all available matkul is less then minimum required SKS"
		return c.JSON(fiber.Map{"error": errorMessage})
	}

	if maxSKS < minSKS {
		c.Status(fiber.StatusBadRequest)
		errorMessage := "Min SKS must be less than or equal to max SKS"
		return c.JSON(fiber.Map{"error": errorMessage})
	}

	bestOptions, IP, SKS := algorithm.FindBestMatkul(filteredMatkul, minSKS, maxSKS)
	
	return c.JSON(fiber.Map{
		"bestOptions": bestOptions,
		"IP": IP,
		"SKS": SKS,
	})

}