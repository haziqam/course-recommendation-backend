package api

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/haziqam/course-scheduler-backend/packages/database"
	"github.com/haziqam/course-scheduler-backend/packages/models"
)

func GetAllJurusan(c *fiber.Ctx) error {
	query := `
		SELECT * 
		FROM jurusan
	`
	rows, err := database.DbInstance.Query(query)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": "Query error"})
	}

	defer rows.Close()

	var jurusanArr []models.Jurusan
	
	for rows.Next() {
		jurusan := new(models.Jurusan)
		err = jurusan.ScanRow(rows);
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{"error": "Error scanning rows"})
		}
		jurusanArr = append(jurusanArr, *jurusan)
	}

	err = rows.Err()
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": "Error iterating rows"})
	}

	return c.JSON(jurusanArr)
}

func addJurusan(c *fiber.Ctx, newJurusan []models.Jurusan) error {
	query := `
		INSERT INTO jurusan(nama_jurusan, nama_fakultas) 
		VALUES ($1, $2)
	`

	for _, jurusan := range newJurusan {
		_, err := database.DbInstance.Exec(query, jurusan.NamaJurusan, jurusan.NamaFakultas)
		if err != nil {
			// log.Fatalln(err)
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{"error": "Query failed"})
		}
	}
	
	c.Status(fiber.StatusCreated)
	return c.JSON(fiber.Map{"message": "Jurusan added successfully"})
}

func AddJurusan(c *fiber.Ctx) error {
	var newJurusan []models.Jurusan
	err := c.BodyParser(&newJurusan)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": "Error parsing request body"})
	}

	return addJurusan(c, newJurusan)
}

func AddJurusanFromFile(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		log.Println("Error parsing form:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error parsing form"})
	}

	files := form.File["Jurusan[]"]
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

	var newJurusan []models.Jurusan
	err = json.Unmarshal(fileContent, &newJurusan)
	if err != nil {
		log.Println("Error unmarshaling file content:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error unmarshaling file content"})
	}

	return addJurusan(c, newJurusan)
}

func RemoveJurusan(c *fiber.Ctx) error {
	namaJurusan := c.Query("jurusan")
	query := `
		DELETE FROM jurusan
		WHERE nama_jurusan = ($1)
	`

	_, err := database.DbInstance.Exec(query, namaJurusan)

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": "Failed to delete jurusan"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}