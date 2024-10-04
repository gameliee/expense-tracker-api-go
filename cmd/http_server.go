package main

import (
	"github.com/labstack/echo/v4"
)

func InitHttp(container *Container) (e *echo.Echo, err error) {
	e = echo.New()
	e.GET("/users/:id", container.UserHandler.GetUser)
	err = nil
	return
}
