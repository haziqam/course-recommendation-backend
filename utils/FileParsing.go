package utils

import (
	"io"

	"github.com/gofiber/fiber/v2"
)

func ParseFileContentFromForm(c *fiber.Ctx, formAttributeName string) ([]byte, error) {
	form, err := c.MultipartForm()
	if err != nil {
		return nil, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error parsing form"})
	}

	files := form.File[formAttributeName]
	if len(files) == 0 {
		return nil, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No file uploaded"})
	}

	file, err := files[0].Open()
	if err != nil {
		return nil, c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error opening file"})
	}
	defer file.Close()

	fileContent, err := io.ReadAll(file)
	if err != nil {
		return nil, c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error reading file"})
	}

	return fileContent, nil
}
