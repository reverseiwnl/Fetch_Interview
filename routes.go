package Fetch_Interview

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/add", addPointsHandler)
	app.Post("/spend", spendPointsHandler)
	app.Get("/balance", balanceHandler)

	err := app.Listen(":8000")
	if err != nil {
		fmt.Print("uh oh! error listening to port!")
	}
}
