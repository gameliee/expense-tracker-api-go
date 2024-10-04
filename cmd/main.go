package main

import (
	"log"
)

func main() {
	echo, err := InitializeContainer()
	if err != nil {
		log.Fatalf("Failed to initialize container: %v", err)
	}

	log.Println("Server starting on :8080, swagger at http://localhost:8080/swagger/index.html")
	log.Fatal(echo.Start(":8080"))
}
