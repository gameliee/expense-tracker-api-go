package cmd

import (
	"log"

	"github.com/labstack/echo/v4"
)

func InitializeApplication() (err error) {
	err = InitializeContainer()
	if err != nil {
		return
	}

	// Initialize database tables
	err = InitializeTables()
	if err != nil {
		return
	}

	// Routes
	err = AttachEndpoints()
	if err != nil {
		return
	}
	return nil
}

func StartApplication() {
	container := GetContainer()
	echo := container.Get((*echo.Echo)(nil)).(*echo.Echo)
	log.Println("Server starting on :8080, swagger at http://localhost:8080/swagger/index.html")
	log.Fatal(echo.Start(":8080"))
}
