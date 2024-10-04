package main

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "gamelieelearn/expense-tracker-api-go/docs"
)

//	@title		Simple Expense Tracker
//	@version	0.1
//	@host		localhost:8080
//	@BasePath	/
func InitHttp(container *Container) (e *echo.Echo, err error) {
	e = echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/users/:id", container.UserHandler.GetUser)
	err = nil
	return
}
