package main

import (
	"log"
)

func main() {
	container, err := InitializeContainer()
	if err != nil {
		log.Fatalf("Failed to initialize container: %v", err)
	}

	// Initialize database tables
	if err := InitializeTables(container.DB); err != nil {
		log.Fatalf("Failed to initialize database tables: %v", err)
	}

	// Routes
	AttachEndpoints(container)

	echo := container.EchoServer
	log.Println("Server starting on :8080, swagger at http://localhost:8080/swagger/index.html")
	log.Fatal(echo.Start(":8080"))
}
