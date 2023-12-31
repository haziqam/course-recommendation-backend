package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/haziqam/course-scheduler-backend/packages/algorithm"
	"github.com/haziqam/course-scheduler-backend/packages/models"
	"github.com/haziqam/course-scheduler-backend/packages/repositories"
	"github.com/haziqam/course-scheduler-backend/packages/utils"
)

var matkulRepo = repositories.GetMatkulRepoInstance()

func GetAllMatkul(c *fiber.Ctx) error {
	matkulArr, err := matkulRepo.GetAllMatkul()
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(matkulArr)
}

func AddMatkul(c *fiber.Ctx) error {
	var newMatkul []models.Matkul
	err := c.BodyParser(&newMatkul)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": "Error parsing request body"})
	}

	err = matkulRepo.AddMatkul(newMatkul)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(fiber.Map{"message": "Matkul added successfully"})
}

func AddMatkulFromFile(c *fiber.Ctx) error {
	fileContent, err := utils.ParseFileContentFromForm(c, "Matkul[]")

	var newMatkul []models.Matkul
	err = json.Unmarshal(fileContent, &newMatkul)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error unmarshaling file content"})
	}

	err = matkulRepo.AddMatkul(newMatkul)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(fiber.Map{"message": "Matkul added successfully"})
}

func RemoveMatkul(c *fiber.Ctx) error {
	namaMatkul := c.Query("matkul")
	namaJurusan := c.Query("jurusan")
	err := matkulRepo.RemoveMatkulByNameAndJurusan(namaMatkul, namaJurusan)

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": "Failed to delete matkul"})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{"message": "Matkul deleted successfully"})
}

type updateMatkulRequestBody struct {
	OldMatkulName           string `json:"oldMatkulName"`
	NewMatkulName           string `json:"newMatkulName"`
	NewMatkulSKS            int    `json:"newMatkulSKS"`
	NewMatkulJurusan        string `json:"newMatkulJurusan"`
	NewMatkulMinSemester    int    `json:"newMatkulMinSemester"`
	NewMatkulPrediksiIndeks string `json:"newMatkulPrediksiIndeks"`
}

func UpdateMatkul(c *fiber.Ctx) error {
	var requestBody updateMatkulRequestBody

	err := c.BodyParser(&requestBody)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": "Error unmarshaling request body"})
	}

	oldMatkulName := requestBody.OldMatkulName

	if oldMatkulName == "" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"error": "Insufficient parameters. Required: oldMatkulName"})
	}

	matkul, err := matkulRepo.GetMatkulByName(oldMatkulName)
	if err != nil || matkul == nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"error": "Failed to find matkul"})
	}

	handlePartialMatkulUpdate(matkul, requestBody)

	err = matkulRepo.UpdateMatkul(oldMatkulName, *matkul)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{"message": "Matkul updated successfully"})
}

func FindMatkul(c *fiber.Ctx) error {
	namaFakultas := c.Query("fakultas")
	currentSemester, err := strconv.Atoi(c.Query("semester"))

	if namaFakultas == "" || err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	filteredMatkul, err := matkulRepo.FilterMatkul(namaFakultas, currentSemester)

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

	filteredMatkul, err := matkulRepo.FilterMatkul(namaFakultas, currentSemester)

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	if len(filteredMatkul) == 0 {
		c.Status(fiber.StatusNotFound)
		errorMessage := fmt.Sprintf("Tidak menemukan matkul dengan nama fakultas = %s dan semester minimum <= %d", namaFakultas, currentSemester)
		return c.JSON(fiber.Map{"error": errorMessage})
	}

	if models.CountTotalSKS(filteredMatkul) < minSKS {
		c.Status(fiber.StatusNotFound)
		errorMessage := "Total SKS matkul yang tersedia lebih sedikit dari SKS minimum yang dapat diambil"
		return c.JSON(fiber.Map{"error": errorMessage})
	}

	if maxSKS < minSKS {
		c.Status(fiber.StatusBadRequest)
		errorMessage := "Min SKS harus lebih kecil atau sama dengan max SKS"
		return c.JSON(fiber.Map{"error": errorMessage})
	}

	bestOptions, IP, SKS := algorithm.FindBestMatkul(filteredMatkul, minSKS, maxSKS)

	return c.JSON(fiber.Map{
		"bestOptions": bestOptions,
		"IP":          IP,
		"SKS":         SKS,
	})
}

func handlePartialMatkulUpdate(
	updatedMatkul *models.Matkul,
	requestBody updateMatkulRequestBody,
) error {
	if requestBody.NewMatkulName != "" {
		updatedMatkul.NamaMatkul = requestBody.NewMatkulName
	}

	if requestBody.NewMatkulSKS != 0 {
		updatedMatkul.SKS = requestBody.NewMatkulSKS
	}

	if requestBody.NewMatkulJurusan != "" {
		updatedMatkul.NamaJurusan = requestBody.NewMatkulJurusan
	}

	if requestBody.NewMatkulMinSemester != 0 {
		updatedMatkul.MinSemester = requestBody.NewMatkulMinSemester
	}

	if requestBody.NewMatkulPrediksiIndeks != "" {
		updatedMatkul.PrediksiIndeks = requestBody.NewMatkulPrediksiIndeks
	}

	return nil
}
