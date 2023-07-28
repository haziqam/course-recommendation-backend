package api

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
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

func AddMatkul(c *fiber.Ctx) error {
	var newMatkul models.Matkul
	err := c.BodyParser(&newMatkul)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": "Error parsing request body"})
	}

	query := `
		INSERT INTO matkul(nama_matkul, sks, nama_jurusan, min_semester, prediksi) 
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err = database.DbInstance.Exec(query, newMatkul.NamaMatkul, newMatkul.SKS, newMatkul.NamaJurusan, 
		newMatkul.MinSemester, newMatkul.PrediksiIndeks)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": "Query failed"})
	}
	
	c.Status(fiber.StatusCreated)
	return c.JSON(fiber.Map{"message": "Matkul added successfully"})
}

func RemoveMatkul(c *fiber.Ctx) error {
	//TODO: implement
}

func FindMatkul(c *fiber.Ctx) error {
	namaFakultas := c.Query("fakultas")
	currentSemester, err := strconv.Atoi(c.Query("semester"))

	if namaFakultas == "" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"error": "Insufficient parameters"})
	}

	query := `
		SELECT nama_matkul, sks, nama_jurusan, min_semester, prediksi
		FROM matkul NATURAL JOIN jurusan
		WHERE nama_fakultas = $1
		AND min_semester <= $2;
	`

	rows, err := database.DbInstance.Query(query, namaFakultas, currentSemester)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": "Query error"})
	}

	defer rows.Close()

	var foundMatkulArr []models.Matkul
	
	for rows.Next() {
		matkul := new(models.Matkul)
		err = matkul.ScanRow(rows);
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{"error": "Error scanning rows"})
		}
		foundMatkulArr = append(foundMatkulArr, *matkul)
	}

	err = rows.Err()
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": "Error iterating rows"})
	}

	return c.JSON(foundMatkulArr)

}