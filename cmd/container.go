package main

import (
	"gamelieelearn/expense-tracker-api-go/config"
	"gamelieelearn/expense-tracker-api-go/internal/http"
	sqliteRepo "gamelieelearn/expense-tracker-api-go/internal/repository/sqlite"
	"gamelieelearn/expense-tracker-api-go/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Container struct {
	Config            *config.Config
	UserRepository    *sqliteRepo.UserRepository
	ExpenseRepository *sqliteRepo.ExpenseRepository
	UserService       *service.UserService
	ExpenseService    *service.ExpenseService
	UserHandler       *http.UserHandler
	ExpenseHandler    *http.ExpenseHandler
	EchoServer        *echo.Echo
	DB                *gorm.DB
}

func NewContainer(
	config *config.Config,
	userRepository *sqliteRepo.UserRepository,
	expenseRepository *sqliteRepo.ExpenseRepository,
	userService *service.UserService,
	expenseService *service.ExpenseService,
	userHandler *http.UserHandler,
	expenseHandler *http.ExpenseHandler,
	echoServer *echo.Echo,
	db *gorm.DB,
) *Container {
	return &Container{
		Config:            config,
		UserRepository:    userRepository,
		ExpenseRepository: expenseRepository,
		UserService:       userService,
		ExpenseService:    expenseService,
		UserHandler:       userHandler,
		ExpenseHandler:    expenseHandler,
		EchoServer:        echoServer,
		DB:                db,
	}
}
