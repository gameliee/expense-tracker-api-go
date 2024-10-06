// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"gamelieelearn/expense-tracker-api-go/config"
	"gamelieelearn/expense-tracker-api-go/internal/http"
	"gamelieelearn/expense-tracker-api-go/internal/repository/sqlite"
	"gamelieelearn/expense-tracker-api-go/service"
)

import (
	_ "gamelieelearn/expense-tracker-api-go/docs"
)

// Injectors from wire.go:

func InitializeContainer() (*Container, error) {
	configConfig := config.NewConfig()
	db, err := InitDB(configConfig)
	if err != nil {
		return nil, err
	}
	userRepository := sqlite.NewUserRepository(db)
	expenseRepository := sqlite.NewExpenseRepository(db)
	userService := service.NewUserService(userRepository)
	expenseService, err := service.NewExpenseService(expenseRepository, userService)
	if err != nil {
		return nil, err
	}
	userHandler := http.NewUserHandler(userService)
	expenseHandler := http.NewExpenseHandler(expenseService)
	echo, err := InitHttp()
	if err != nil {
		return nil, err
	}
	container := NewContainer(configConfig, userRepository, expenseRepository, userService, expenseService, userHandler, expenseHandler, echo, db)
	return container, nil
}
