package api

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/haziqam/course-scheduler-backend/packages/database"
	"github.com/haziqam/course-scheduler-backend/packages/models"
)

func GetAllFakultas(c *fiber.Ctx) error {
	query := `
		SELECT * 
		FROM fakultas
	`
	rows, err := database.DbInstance.Query(query)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": "Query error"})
	}

	defer rows.Close()

	var fakultasArr []models.Fakultas
	
	for rows.Next() {
		fakultas := new(models.Fakultas)
		err = fakultas.ScanRow(rows)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{"error": "Error scanning rows"})
		}
		fakultasArr = append(fakultasArr, *fakultas)
	}

	err = rows.Err()
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": "Error iterating rows"})
	}

	return c.JSON(fakultasArr)
}

func addFakultas(c *fiber.Ctx, newFakultas []models.Fakultas) error {
	query := `
		INSERT INTO fakultas(nama_fakultas)
		VALUES ($1)
	`

	for _, fakultas := range newFakultas {
		_, err := database.DbInstance.Exec(query, fakultas.NamaFakultas)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{"error": "Query failed"})
		}
	}
	
	c.Status(fiber.StatusCreated)
	return c.JSON(fiber.Map{"message": "Fakultas added successfully"})
}

func AddFakultas(c *fiber.Ctx) error {
	var newFakultas []models.Fakultas
	err := c.BodyParser(&newFakultas)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": "Error parsing request body"})
	}

	return addFakultas(c, newFakultas)
}

func AddFakultasFromFile(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		log.Println("Error parsing form:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error parsing form"})
	}

	files := form.File["Fakultas[]"]
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

	var newFakultas []models.Fakultas
	err = json.Unmarshal(fileContent, &newFakultas)
	if err != nil {
		log.Println("Error unmarshaling file content:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error unmarshaling file content"})
	}

	return addFakultas(c, newFakultas)
}

func RemoveFakultas(c *fiber.Ctx) error {
	namaFakultas := c.Query("fakultas")
	query := `
		DELETE FROM fakultas
		WHERE nama_fakultas = ($1)
	`

	_, err := database.DbInstance.Exec(query, namaFakultas)

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": "Failed to delete fakultas"})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{"message": "Fakultas deleted successfully"})
}