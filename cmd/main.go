package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/haziqam/course-scheduler-backend/packages/database"
)

func main() {
    database.ConnectDB()
    app := fiber.New()
    app.Use(cors.New())
    initRoutes(app)
    app.Listen(":5000")
}