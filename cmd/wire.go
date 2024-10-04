//go:build wireinject
// +build wireinject

package main

import (
	"gamelieelearn/expense-tracker-api-go/config"
	"gamelieelearn/expense-tracker-api-go/internal/http"
	sqliteRepo "gamelieelearn/expense-tracker-api-go/internal/repository/sqlite"
	"gamelieelearn/expense-tracker-api-go/service"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

func InitializeContainer() (*echo.Echo, error) {
	wire.Build(
		wire.Bind(new(service.UserRepositoryInt), new(*sqliteRepo.UserRepository)),
		wire.Bind(new(service.ExpenseRepositoryInt), new(*sqliteRepo.ExpenseRepository)),
		wire.Bind(new(http.UserServiceInt), new(*service.UserService)),
		wire.Bind(new(http.ExpenseServiceInt), new(*service.ExpenseService)),
		config.NewConfig,
		sqliteRepo.NewUserRepository,
		sqliteRepo.NewExpenseRepository,
		service.NewUserService,
		service.NewExpenseService,
		http.NewUserHandler,
		http.NewExpenseHandler,
		NewContainer,
		InitDB,
		InitHttp,
	)
	return &echo.Echo{}, nil
}
