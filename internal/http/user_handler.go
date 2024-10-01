package http

import (
	"context"
	"net/http"

	"gamelieelearn/expense-tracker-api-go/domain"
)

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

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	// Implement the handler
}
