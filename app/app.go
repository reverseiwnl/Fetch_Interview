package app

import (
	"github.com/gofiber/fiber/v2"
)

func SetupApp() *fiber.App {
	app := fiber.New()
	SetupRoutes(app)
	return app
}
