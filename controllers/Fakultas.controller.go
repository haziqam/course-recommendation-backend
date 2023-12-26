package controllers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/haziqam/course-scheduler-backend/packages/models"
	"github.com/haziqam/course-scheduler-backend/packages/repositories"
	"github.com/haziqam/course-scheduler-backend/packages/utils"
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
	fileContent, err := utils.ParseFileContentFromForm(c, "Fakultas[]")
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	var newFakultas []models.Fakultas
	err = json.Unmarshal(fileContent, &newFakultas)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": "Error unmarshaling file content"})
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

	fakultas, err := fakultasRepo.GetFakultasByName(namaFakultas)
	if err != nil || fakultas == nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": "Failed to find fakultas"})
	}

	err = fakultasRepo.RemoveFakultasByName(namaFakultas)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": "Failed to delete fakultas"})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{"message": "Fakultas deleted successfully"})
}

func UpdateFakultas(c *fiber.Ctx) error {
	var requestBody = c.Body()
	var requestBodyMap map[string]string

	err := json.Unmarshal(requestBody, &requestBodyMap)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": "Error unmarshaling file content"})
	}

	oldFakultasName := requestBodyMap["oldFakultasName"]
	newFakultasName := requestBodyMap["newFakultasName"]

	if oldFakultasName == "" || newFakultasName == "" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"error": "Insufficient parameters. Required: oldFakultasName, newFakultasName"})
	}

	err = fakultasRepo.UpdateFakultas(oldFakultasName, newFakultasName)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{"message": "Fakultas updated successfully"})
}
