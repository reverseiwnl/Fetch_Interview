package main

import (
	app2 "Fetch_Interview/app"
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestSolution(t *testing.T) {
	app := app2.SetupApp()

	transactions := []map[string]interface{}{
		{"payer": "DANNON", "points": 300, "timestamp": "2022-10-31T10:00:00Z"},
		{"payer": "UNILEVER", "points": 200, "timestamp": "2022-10-31T11:00:00Z"},
		{"payer": "DANNON", "points": -200, "timestamp": "2022-10-31T15:00:00Z"},
		{"payer": "MILLER COORS", "points": 10000, "timestamp": "2022-11-01T14:00:00Z"},
		{"payer": "DANNON", "points": 1000, "timestamp": "2022-11-02T14:00:00Z"},
	}

	for _, transaction := range transactions {
		resp, err := sendRequest(app, "POST", "/add", transaction)
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode, "Failed to add transaction")
	}

	spendRequest := map[string]interface{}{
		"points": 5000,
	}
	resp, err := sendRequest(app, "POST", "/spend", spendRequest)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode, "Failed to spend points")

	resp, err = sendRequest(app, "GET", "/balance", nil)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode, "Failed to retrieve balance")

	var balance map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&balance)
	assert.NoError(t, err, "Failed to parse balance response")

	expectedBalance := map[string]interface{}{
		"DANNON":       float64(1000),
		"UNILEVER":     float64(0),
		"MILLER COORS": float64(5300),
	}
	assert.Equal(t, expectedBalance, balance, "Balance results mismatch")
}

func sendRequest(app *fiber.App, method, url string, body interface{}) (*http.Response, error) {
	var reqBody []byte
	var err error

	if body != nil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return app.Test(req)
}
