package main

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "gamelieelearn/expense-tracker-api-go/docs"
)

// @title		Simple Expense Tracker
// @version	0.1
// @host		localhost:8080
// @BasePath	/
func InitHttp(container *Container) (e *echo.Echo, err error) {
	e = echo.New()

	// Swagger documentation
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// User routes
	e.POST("/users", container.UserHandler.CreateUser)
	e.GET("/users/:id", container.UserHandler.GetUser)
	e.PUT("/users/:id", container.UserHandler.UpdateUser)
	e.DELETE("/users/:id", container.UserHandler.DeleteUser)
	e.GET("/users", container.UserHandler.ListUsers)

	// Expense routes
	e.POST("/expenses", container.ExpenseHandler.CreateExpense)
	e.GET("/expenses/:id", container.ExpenseHandler.GetExpense)
	e.PUT("/expenses/:id", container.ExpenseHandler.UpdateExpense)
	e.DELETE("/expenses/:id", container.ExpenseHandler.DeleteExpense)
	e.GET("/expenses", container.ExpenseHandler.ListExpenses)
	e.GET("/users/:user_id/expenses", container.ExpenseHandler.GetExpensesByUserID)

	err = nil
	return
}
