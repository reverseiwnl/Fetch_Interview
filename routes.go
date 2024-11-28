package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Post("/add", addPointsHandler)
	app.Post("/spend", spendPointsHandler)
	app.Get("/balance", balanceHandler)
}

func main() {
	app := fiber.New()
	setupRoutes(app)

	// Start the server
	log.Println("Starting server on http://localhost:8000")
	if err := app.Listen(":8000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
