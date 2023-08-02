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
	app.Delete("/fakultas", api.RemoveFakultas)
	app.Post("/fakultas/addFromFile", api.AddFakultasFromFile)

	app.Get("/jurusan", api.GetAllJurusan)
	app.Post("/jurusan", api.AddJurusan)
	app.Delete("/jurusan", api.RemoveJurusan)
	app.Post("/jurusan/addFromFile", api.AddJurusanFromFile)

	app.Get("/matkul", api.GetAllMatkul)
	app.Post("/matkul", api.AddMatkul)
	app.Delete("/matkul", api.RemoveMatkul)
	app.Post("/matkul/addFromFile", api.AddMatkulFromFile)
	app.Get("/matkul/find", api.FindMatkul)
	app.Get("/matkul/find/bestOptions", api.FindBestOptions)
}