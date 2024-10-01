//go:build wireinject
// +build wireinject

package main

import (
	"gamelieelearn/expense-tracker-api-go/config"
	"gamelieelearn/expense-tracker-api-go/internal/http"
	sqliteRepo "gamelieelearn/expense-tracker-api-go/internal/repository/sqlite"
	"gamelieelearn/expense-tracker-api-go/service"

	"github.com/google/wire"
)

func InitializeContainer() (Container, error) {
	wire.Build(
		config.NewConfig,
		sqliteRepo.NewUserRepository,
		wire.Bind(new(service.UserRepository), new(sqliteRepo.UserRepository)),
		sqliteRepo.NewExpenseRepository,
		wire.Bind(new(service.ExpenseRepository), new(sqliteRepo.ExpenseRepository)),
		service.NewUserService,
		service.NewExpenseService,
		wire.Bind(new(http.UserService), new(service.UserService)),
		wire.Bind(new(http.ExpenseService), new(service.ExpenseService)),
		http.NewUserHandler,
		http.NewExpenseHandler,
		NewContainer,
	)
	return Container{}, nil
}
