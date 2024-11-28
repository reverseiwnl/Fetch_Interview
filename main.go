package main

import (
	app2 "Fetch_Interview/app"
	"log"
)

func main() {
	app := app2.SetupApp()
	log.Println("Starting server on http://localhost:8000")
	if err := app.Listen(":8000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
