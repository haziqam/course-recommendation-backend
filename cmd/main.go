package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/haziqam/course-scheduler-backend/packages/database"
)

func main() {
    database.ConnectDB()

    app := fiber.New()

    initRoutes(app)

    app.Listen(":3000")
}