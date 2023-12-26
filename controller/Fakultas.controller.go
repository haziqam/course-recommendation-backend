package controller

import (
	"encoding/json"
	"io"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/haziqam/course-scheduler-backend/packages/models"
	"github.com/haziqam/course-scheduler-backend/packages/repositories"
)

var fakultasRepo = repositories.GetFakultasRepoInstance()

func GetAllFakultas(c *fiber.Ctx) error {
	fakultasArr, err := fakultasRepo.GetAllFakultas()
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fakultasArr)
}

func AddFakultas(c *fiber.Ctx) error {
	var newFakultas []models.Fakultas
	err := c.BodyParser(&newFakultas)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": "Error parsing request body"})
	}

	err = fakultasRepo.AddFakultas(newFakultas)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(fiber.Map{"message": "Fakultas added successfully"})
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

	fileContent, err := io.ReadAll(file)
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

	err = fakultasRepo.AddFakultas(newFakultas)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(fiber.Map{"message": "Fakultas added successfully"})
}

func RemoveFakultas(c *fiber.Ctx) error {
	namaFakultas := c.Query("fakultas")
	err := fakultasRepo.RemoveFakultasByName(namaFakultas)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": "Failed to delete fakultas"})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{"message": "Fakultas deleted successfully"})
}
