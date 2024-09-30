package http

import (
	"context"
	"net/http"

	"gamelieelearn/expense-tracker-api-go/domain"
)

type UserService interface {
	Store(ctx context.Context, u *domain.User) (err error)
	GetByID(ctx context.Context, id int64) (res domain.User, err error)
}

type UserHandler struct {
	UserService UserService
}

func NewUserHandler(userService UserService) UserHandler {
	return UserHandler{
		UserService: userService,
	}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	// Implement the handler
}
