package app

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	app.Post("/add", AddPointsHandler)
	app.Post("/spend", SpendPointsHandler)
	app.Get("/balance", BalanceHandler)
}
