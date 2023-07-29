package api

import (
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

func AddJurusan(c *fiber.Ctx) error {
	var newJurusan []models.Jurusan
	err := c.BodyParser(&newJurusan)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": "Error parsing request body"})
	}

	query := `
		INSERT INTO jurusan(nama_jurusan, nama_fakultas) 
		VALUES ($1, $2)
	`

	for _, jurusan := range newJurusan {
		_, err = database.DbInstance.Exec(query, jurusan.NamaJurusan, jurusan.NamaFakultas)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{"error": "Query failed"})
		}
	}
	
	c.Status(fiber.StatusCreated)
	return c.JSON(fiber.Map{"message": "Jurusan added successfully"})
}

// func RemoveJurusan(c *fiber.Ctx) error {
// 	//TODO: implement
// }