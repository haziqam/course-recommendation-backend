package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/haziqam/course-scheduler-backend/packages/api"
)

func initRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, masseh")
	})

	app.Get("/fakultas", api.GetAllFakultas)
	app.Post("/fakultas", api.AddFakultas)

	app.Get("/jurusan", api.GetAllJurusan)
	app.Post("/jurusan", api.AddJurusan)

	app.Get("/matkul", api.GetAllMatkul)
	app.Post("/matkul", api.AddMatkul)
	app.Get("/matkul/find", api.FindMatkul)
}