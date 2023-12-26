package controllers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/haziqam/course-scheduler-backend/packages/models"
	"github.com/haziqam/course-scheduler-backend/packages/repositories"
	"github.com/haziqam/course-scheduler-backend/packages/utils"
)

var jurusanRepo = repositories.GetJurusanRepoInstance()

func GetAllJurusan(c *fiber.Ctx) error {
	jurusanArr, err := jurusanRepo.GetAllJurusan()
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(jurusanArr)
}

func AddJurusan(c *fiber.Ctx) error {
	var newJurusan []models.Jurusan
	err := c.BodyParser(&newJurusan)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": "Error parsing request body"})
	}

	err = jurusanRepo.AddJurusan(newJurusan)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(fiber.Map{"message": "Jurusan added successfully"})
}

func AddJurusanFromFile(c *fiber.Ctx) error {
	fileContent, err := utils.ParseFileContentFromForm(c, "Jurusan[]")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	var newJurusan []models.Jurusan
	err = json.Unmarshal(fileContent, &newJurusan)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error unmarshaling file content"})
	}

	err = jurusanRepo.AddJurusan(newJurusan)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(fiber.Map{"message": "Jurusan added successfully"})
}

func RemoveJurusan(c *fiber.Ctx) error {
	namaJurusan := c.Query("jurusan")
	err := jurusanRepo.RemoveJurusanByName(namaJurusan)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": "Failed to delete jurusan"})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{"message": "Jurusan deleted successfully"})
}

func UpdateJurusan(c *fiber.Ctx) error {
	var requestBody = c.Body()
	var requestBodyMap map[string]string

	err := json.Unmarshal(requestBody, &requestBodyMap)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": "Error unmarshaling file content"})
	}

	oldJurusanName := requestBodyMap["oldJurusanName"]
	newJurusanName := requestBodyMap["newJurusanName"]
	newJurusanFakultas := requestBodyMap["newJurusanFakultas"]

	jurusan, err := jurusanRepo.GetJurusanByName(oldJurusanName)

	if err != nil || jurusan == nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"error": "Failed to find jurusan"})
	}

	if oldJurusanName == "" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"error": "Insufficient parameters. Required: oldJurusanName"})
	}

	if newJurusanName != "" {
		jurusan.NamaJurusan = newJurusanName
	}

	if newJurusanFakultas != "" {
		jurusan.NamaFakultas = newJurusanFakultas
	}

	err = jurusanRepo.UpdateJurusan(oldJurusanName, *jurusan)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": "Failed to update jurusan"})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{"message": "Jurusan updated successfully"})
}
