package main

import (
	"gamelieelearn/expense-tracker-api-go/config"
	"gamelieelearn/expense-tracker-api-go/internal/http"
	sqliteRepo "gamelieelearn/expense-tracker-api-go/internal/repository/sqlite"
	"gamelieelearn/expense-tracker-api-go/service"
)

type Container struct {
	Config            *config.Config
	UserRepository    *sqliteRepo.UserRepository
	ExpenseRepository *sqliteRepo.ExpenseRepository
	UserService       *service.UserService
	ExpenseService    *service.ExpenseService
	UserHandler       *http.UserHandler
	ExpenseHandler    *http.ExpenseHandler
}

func NewContainer(
	config *config.Config,
	userRepository *sqliteRepo.UserRepository,
	expenseRepository *sqliteRepo.ExpenseRepository,
	userService *service.UserService,
	expenseService *service.ExpenseService,
	userHandler *http.UserHandler,
	expenseHandler *http.ExpenseHandler,
) *Container {
	return &Container{
		Config:            config,
		UserRepository:    userRepository,
		ExpenseRepository: expenseRepository,
		UserService:       userService,
		ExpenseService:    expenseService,
		UserHandler:       userHandler,
		ExpenseHandler:    expenseHandler,
	}
}
