package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/haziqam/course-scheduler-backend/packages/api"
)

func initRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, masseh")
	})

	app.Get("/fakultas", api.GetFakultas)

	app.Post("/fakultas", api.AddFakultas)
}