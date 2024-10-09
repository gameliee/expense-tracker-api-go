package main

import (
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	err := InitializeContainer()
	if err != nil {
		log.Fatalf("Failed to initialize container: %v", err)
	}

	// Initialize database tables
	err = InitializeTables()
	if err != nil {
		log.Fatalf("Failed to initialize database tables: %v", err)
	}

	// Routes
	err = AttachEndpoints()
	if err != nil {
		log.Fatalf("Failed to attach endpoints: %v", err)
	}

	echo := container.Get((*echo.Echo)(nil)).(*echo.Echo)
	log.Println("Server starting on :8080, swagger at http://localhost:8080/swagger/index.html")
	log.Fatal(echo.Start(":8080"))
}
