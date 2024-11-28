package main

import (
	"github.com/gofiber/fiber/v2"
	"sort"
	"sync"
)

type SpendResponse struct {
	Payer  string `json:"payer"`
	Points int    `json:"points"`
}

type SpendRequest struct {
	Points int `json:"points"`
}

var (
	points       = make(map[string]int)
	transactions []Transaction
	mutex        sync.Mutex
)

func addPointsHandler(c *fiber.Ctx) error {
	var transaction Transaction
	err := c.BodyParser(&transaction)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	mutex.Lock()
	defer mutex.Unlock()

	points[transaction.Payer] += transaction.Points
	transactions = append(transactions, transaction)

	sort.Sort(TransactionsByTimestamp(transactions))

	return c.SendStatus(fiber.StatusOK)
}

func spendPointsHandler(c *fiber.Ctx) error {
	var spendRequest SpendRequest
	err := c.BodyParser(&spendRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	mutex.Lock()
	defer mutex.Unlock()

	totalPointsToSpend := spendRequest.Points
	var response []SpendResponse

	if totalPointsToSpend > totalPoints() {
		return c.Status(fiber.StatusBadRequest).SendString("Not enough points")
	}

	for totalPointsToSpend > 0 && len(transactions) > 0 {
		t := &transactions[0]
		spend := min(totalPointsToSpend, t.Points)

		if points[t.Payer]-spend < 0 {
			break
		}

		points[t.Payer] -= spend
		t.Points -= spend
		totalPointsToSpend -= spend

		response = append(response, SpendResponse{
			Payer:  t.Payer,
			Points: -spend,
		})

		if t.Points == 0 {
			transactions = transactions[1:]
		}
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func totalPoints() int {
	sum := 0
	for _, pts := range points {
		sum += pts
	}
	return sum
}

func balanceHandler(c *fiber.Ctx) error {
	mutex.Lock()
	defer mutex.Unlock()

	return c.Status(fiber.StatusOK).JSON(points)
}
