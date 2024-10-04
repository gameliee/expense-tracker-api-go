package http

import (
	"context"
	"net/http"
	"strconv"

	"gamelieelearn/expense-tracker-api-go/domain"

	"github.com/labstack/echo/v4"
)

type ResponseError struct {
	Message string `json:"message"`
}

type UserServiceInt interface {
	Store(ctx context.Context, u *domain.User) (err error)
	GetByID(ctx context.Context, id int64) (res domain.User, err error)
}

type UserHandler struct {
	UserService UserServiceInt
}

func NewUserHandler(userService UserServiceInt) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (h *UserHandler) GetUser(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	id := int64(idP)
	ctx := c.Request().Context()

	user, err := h.UserService.GetByID(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, user)
}
