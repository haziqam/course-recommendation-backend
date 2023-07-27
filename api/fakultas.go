package api

import (
	"github.com/gofiber/fiber/v2"
)

func GetFakultas(c *fiber.Ctx) error {
	// TODO: menquery untuk SELECT * FROM fakultas, simpen di array, return array
	return c.SendString("nich fakultas")
}

func AddFakultas(c *fiber.Ctx) error {
	// TODO: parse body, query untuk INSERT INTO FAKULTAS VALUES(...)
	return c.SendString("hehhe")
}