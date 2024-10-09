package main

import (
	"gamelieelearn/expense-tracker-api-go/config"
	"gamelieelearn/expense-tracker-api-go/internal/http"
	"gamelieelearn/expense-tracker-api-go/internal/repository/sqlite"
	"gamelieelearn/expense-tracker-api-go/service"
	"gamelieelearn/expense-tracker-api-go/tools"
	"sync"
)

var (
	container *tools.Container
	once      sync.Once
)

func GetContainer() *tools.Container {
	once.Do(func() {
		container = tools.NewContainer()
	})
	return container
}

func InitializeContainer() error {
	c := GetContainer()
	config := config.NewConfig()
	db, err := InitDB(config)
	if err != nil {
		return err
	}
	c.RegisterInstance(config)
	c.RegisterInstance(db)
	c.RegisterInstance(&sqlite.ExpenseRepository{})
	c.RegisterInstance(&sqlite.UserRepository{})
	c.RegisterInstance(&service.UserService{})
	c.RegisterInstance(&service.ExpenseService{})
	c.RegisterInstance(&http.UserHandler{})
	c.RegisterInstance(&http.ExpenseHandler{})
	echo, err := InitHttp()
	if err != nil {
		return err
	}
	c.RegisterInstance(echo)
	return c.Build()
}
