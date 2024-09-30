package main

import (
	"log"
	"net/http"
)

func main() {
	container, err := InitializeContainer()
	if err != nil {
		log.Fatalf("Failed to initialize container: %v", err)
	}

	// Set up HTTP routes
	http.HandleFunc("/user", container.UserHandler.GetUser)
	http.HandleFunc("/expense", container.ExpenseHandler.GetExpense)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
