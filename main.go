package main

import (
	"mini-project/database"
	"mini-project/database/migration"
	"mini-project/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.DatabaseInit()
	migration.RunMigration()

	app := fiber.New()

	route.RouteInit(app)

	app.Listen(":3030")
}