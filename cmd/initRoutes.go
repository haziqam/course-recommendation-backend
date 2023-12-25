package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/haziqam/course-scheduler-backend/packages/controller"
)

func initRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, masseh")
	})

	app.Get("/fakultas", controller.GetAllFakultas)
	app.Post("/fakultas", controller.AddFakultas)
	app.Delete("/fakultas", controller.RemoveFakultas)
	app.Post("/fakultas/addFromFile", controller.AddFakultasFromFile)

	app.Get("/jurusan", controller.GetAllJurusan)
	app.Post("/jurusan", controller.AddJurusan)
	app.Delete("/jurusan", controller.RemoveJurusan)
	app.Post("/jurusan/addFromFile", controller.AddJurusanFromFile)

	app.Get("/matkul", controller.GetAllMatkul)
	app.Post("/matkul", controller.AddMatkul)
	app.Delete("/matkul", controller.RemoveMatkul)
	app.Post("/matkul/addFromFile", controller.AddMatkulFromFile)
	app.Get("/matkul/find", controller.FindMatkul)
	app.Get("/matkul/find/bestOptions", controller.FindBestOptions)
}
