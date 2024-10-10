package cmd

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "gamelieelearn/expense-tracker-api-go/docs"
	"gamelieelearn/expense-tracker-api-go/internal/http"
)

// @title		Simple Expense Tracker
// @version	0.1
// @host		localhost:8080
// @BasePath	/
func AttachEndpoints() error {
	container := GetContainer()
	e := container.Get((*echo.Echo)(nil)).(*echo.Echo)
	user_handler := container.Get((*http.UserHandler)(nil)).(*http.UserHandler)
	expense_handler := container.Get((*http.ExpenseHandler)(nil)).(*http.ExpenseHandler)

	// Swagger documentation
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// User routes
	e.POST("/users", user_handler.CreateUser)
	e.GET("/users/:id", user_handler.GetUser)
	e.PUT("/users/:id", user_handler.UpdateUser)
	e.DELETE("/users/:id", user_handler.DeleteUser)
	e.GET("/users", user_handler.ListUsers)

	// Expense routes
	e.POST("/expenses", expense_handler.CreateExpense)
	e.GET("/expenses/:id", expense_handler.GetExpense)
	e.PUT("/expenses/:id", expense_handler.UpdateExpense)
	e.DELETE("/expenses/:id", expense_handler.DeleteExpense)
	e.GET("/expenses", expense_handler.ListExpenses)
	e.GET("/users/:user_id/expenses", expense_handler.GetExpensesByUserID)
	return nil
}

func InitHttp() (e *echo.Echo, err error) {
	e = echo.New()
	err = nil
	return
}
