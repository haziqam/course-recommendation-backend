package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/haziqam/course-scheduler-backend/packages/controllers"
)

func InitRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, masseh")
	})

	app.Get("/fakultas", controllers.GetAllFakultas)
	app.Post("/fakultas", controllers.AddFakultas)
	app.Delete("/fakultas", controllers.RemoveFakultas)
	app.Put("/fakultas", controllers.UpdateFakultas)
	app.Post("/fakultas/addFromFile", controllers.AddFakultasFromFile)

	app.Get("/jurusan", controllers.GetAllJurusan)
	app.Post("/jurusan", controllers.AddJurusan)
	app.Delete("/jurusan", controllers.RemoveJurusan)
	app.Patch("/jurusan", controllers.UpdateJurusan)
	app.Post("/jurusan/addFromFile", controllers.AddJurusanFromFile)

	app.Get("/matkul", controllers.GetAllMatkul)
	app.Post("/matkul", controllers.AddMatkul)
	app.Delete("/matkul", controllers.RemoveMatkul)
	app.Post("/matkul/addFromFile", controllers.AddMatkulFromFile)
	app.Get("/matkul/find", controllers.FindMatkul)
	app.Get("/matkul/find/bestOptions", controllers.FindBestOptions)
}
