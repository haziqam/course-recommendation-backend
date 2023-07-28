package api

import (
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
//

func AddFakultas(c *fiber.Ctx) error {
	var newFakultas models.Fakultas
	err := c.BodyParser(&newFakultas)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": "Error parsing request body"})
	}

	query := `
		INSERT INTO fakultas(nama_fakultas)
		VALUES ($1)
	`
	_, err = database.DbInstance.Exec(query, newFakultas.NamaFakultas)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": "Query failed"})
	}
	
	c.Status(fiber.StatusCreated)
	return c.JSON(fiber.Map{"message": "Fakultas added successfully"})
}

func RemoveFakultas(c *fiber.Ctx) error {
	//TODO: implement
}