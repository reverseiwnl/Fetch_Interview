package Fetch_Interview

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	go func() {
		app := setupApp()
		err := app.Listen(":8000")
		if err != nil {
			panic(err)
		}
	}()

	client := resty.New()

	transactions := []map[string]interface{}{
		{"payer": "DANNON", "points": 300, "timestamp": "2022-10-31T10:00:00Z"},
		{"payer": "UNILEVER", "points": 200, "timestamp": "2022-10-31T11:00:00Z"},
		{"payer": "DANNON", "points": -200, "timestamp": "2022-10-31T15:00:00Z"},
		{"payer": "MILLER COORS", "points": 10000, "timestamp": "2022-11-01T14:00:00Z"},
		{"payer": "DANNON", "points": 1000, "timestamp": "2022-11-02T14:00:00Z"},
	}

	for _, transaction := range transactions {
		resp, err := client.R().
			SetBody(transaction).
			Post("http://localhost:8000/add")

		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode(), "Failed to add transaction")
	}

	// Spend points
	spendRequest := map[string]interface{}{
		"points": 5000,
	}
	spendResponse, err := client.R().
		SetBody(spendRequest).
		Post("http://localhost:8000/spend")

	assert.NoError(t, err)
	assert.Equal(t, 200, spendResponse.StatusCode(), "Failed to spend points")

	// Check balance
	balanceResponse, err := client.R().
		Get("http://localhost:8000/balance")

	assert.NoError(t, err)
	assert.Equal(t, 200, balanceResponse.StatusCode(), "Failed to get balance")

	// Parse and validate balance response
	var balance map[string]interface{}
	err = json.Unmarshal(balanceResponse.Body(), &balance)
	assert.NoError(t, err, "Failed to parse balance response")

	expectedBalance := map[string]interface{}{
		"DANNON":       (float64)(1000),
		"UNILEVER":     (float64)(0),
		"MILLER COORS": (float64)(5300),
	}
	assert.Equal(t, expectedBalance, balance, "Balance results mismatch")
}

func setupApp() *fiber.App {
	// Set up the app with the same routes as in routes.go
	app := fiber.New()

	app.Post("/add", addPointsHandler)
	app.Post("/spend", spendPointsHandler)
	app.Get("/balance", balanceHandler)

	return app
}
