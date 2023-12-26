package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/haziqam/course-scheduler-backend/packages/database"
	"github.com/haziqam/course-scheduler-backend/packages/routes"
)

func main() {
	database.ConnectDB()
	app := fiber.New()
	app.Use(cors.New())
	routes.InitRoutes(app)
	app.Listen(":5000")
}
