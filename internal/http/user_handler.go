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

//go:generate mockery --name=UserServiceInt --struct=UserService
type UserServiceInt interface {
	Store(ctx context.Context, u *domain.User) error
	GetByID(ctx context.Context, id int64) (domain.User, error)
	Update(ctx context.Context, u *domain.User) error
	Delete(ctx context.Context, id int64) error
	GetAll(ctx context.Context) ([]domain.User, error)
}

type UserHandler struct {
	UserService UserServiceInt
}

func NewUserHandler(userService UserServiceInt) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

// CreateUser godoc
//
//	@Summary		Create a new user
//	@Description	create a new user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		domain.User	true	"User object"
//	@Success		201		{object}	domain.User
//	@Failure		400		{object}	ResponseError
//	@Router			/users [post]
func (h *UserHandler) CreateUser(c echo.Context) error {
	var user domain.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: "invalid request"})
	}

	ctx := c.Request().Context()
	if err := h.UserService.Store(ctx, &user); err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, user)
}

// GetUser godoc
//
//	@Summary		Get a user by ID
//	@Description	get user by ID
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	domain.User
//	@Failure		400	{object}	ResponseError
//	@Failure		404	{object}	ResponseError
//	@Router			/users/{id} [get]
func (h *UserHandler) GetUser(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: "invalid id"})
	}
	id := int64(idP)
	ctx := c.Request().Context()

	user, err := h.UserService.GetByID(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, user)
}

// UpdateUser godoc
//
//	@Summary		Update a user
//	@Description	update a user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int			true	"User ID"
//	@Param			user	body		domain.User	true	"User object"
//	@Success		200		{object}	domain.User
//	@Failure		400		{object}	ResponseError
//	@Failure		404		{object}	ResponseError
//	@Router			/users/{id} [put]
func (h *UserHandler) UpdateUser(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: "invalid id"})
	}
	id := int64(idP)

	var user domain.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: "invalid request"})
	}
	user.ID = id

	ctx := c.Request().Context()
	if err := h.UserService.Update(ctx, &user); err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}

// DeleteUser godoc
//
//	@Summary		Delete a user
//	@Description	delete a user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int		true	"User ID"
//	@Success		204	{string}	string	"No Content"
//	@Failure		400	{object}	ResponseError
//	@Failure		404	{object}	ResponseError
//	@Router			/users/{id} [delete]
func (h *UserHandler) DeleteUser(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: "invalid id"})
	}
	id := int64(idP)

	ctx := c.Request().Context()
	if err := h.UserService.Delete(ctx, id); err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

// ListUsers godoc
//
//	@Summary		List all users
//	@Description	get all users
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		domain.User
//	@Failure		500	{object}	ResponseError
//	@Router			/users [get]
func (h *UserHandler) ListUsers(c echo.Context) error {
	ctx := c.Request().Context()
	users, err := h.UserService.GetAll(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, users)
}
