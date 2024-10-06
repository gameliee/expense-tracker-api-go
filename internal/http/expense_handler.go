package http

import (
	"context"
	"gamelieelearn/expense-tracker-api-go/domain"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

//go:generate mockery --name=ExpenseServiceInt --structname=ExpenseService
type ExpenseServiceInt interface {
	Store(ctx context.Context, expense *domain.Expense) error
	GetByID(ctx context.Context, id int64) (domain.Expense, error)
	Update(ctx context.Context, expense *domain.Expense) error
	Delete(ctx context.Context, id int64) error
	GetAll(ctx context.Context) ([]domain.Expense, error)
	GetByUserID(ctx context.Context, userID int64) ([]domain.Expense, error)
}

type ExpenseHandler struct {
	ExpenseService ExpenseServiceInt
}

func NewExpenseHandler(expenseService ExpenseServiceInt) *ExpenseHandler {
	return &ExpenseHandler{
		ExpenseService: expenseService,
	}
}

// CreateExpense godoc
//
//	@Summary		Create a new expense
//	@Description	create a new expense
//	@Tags			expenses
//	@Accept			json
//	@Produce		json
//	@Param			expense	body		domain.Expense	true	"Expense object"
//	@Success		201		{object}	domain.Expense
//	@Failure		400		{object}	ResponseError
//	@Router			/expenses [post]
func (h *ExpenseHandler) CreateExpense(c echo.Context) error {
	var expense domain.Expense
	if err := c.Bind(&expense); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: "invalid request"})
	}

	ctx := c.Request().Context()
	if err := h.ExpenseService.Store(ctx, &expense); err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, expense)
}

// GetExpense godoc
//
//	@Summary		Get an expense by ID
//	@Description	get expense by ID
//	@Tags			expenses
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Expense ID"
//	@Success		200	{object}	domain.Expense
//	@Failure		400	{object}	ResponseError
//	@Failure		404	{object}	ResponseError
//	@Router			/expenses/{id} [get]
func (h *ExpenseHandler) GetExpense(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: "invalid id"})
	}

	ctx := c.Request().Context()
	expense, err := h.ExpenseService.GetByID(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, expense)
}

// UpdateExpense godoc
//
//	@Summary		Update an expense
//	@Description	update an expense
//	@Tags			expenses
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int				true	"Expense ID"
//	@Param			expense	body		domain.Expense	true	"Expense object"
//	@Success		200		{object}	domain.Expense
//	@Failure		400		{object}	ResponseError
//	@Failure		404		{object}	ResponseError
//	@Router			/expenses/{id} [put]
func (h *ExpenseHandler) UpdateExpense(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: "invalid id"})
	}

	var expense domain.Expense
	if err := c.Bind(&expense); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: "invalid request"})
	}
	expense.ID = id

	ctx := c.Request().Context()
	if err := h.ExpenseService.Update(ctx, &expense); err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, expense)
}

// DeleteExpense godoc
//
//	@Summary		Delete an expense
//	@Description	delete an expense
//	@Tags			expenses
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int		true	"Expense ID"
//	@Success		204	{string}	string	"No Content"
//	@Failure		400	{object}	ResponseError
//	@Failure		404	{object}	ResponseError
//	@Router			/expenses/{id} [delete]
func (h *ExpenseHandler) DeleteExpense(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: "invalid id"})
	}

	ctx := c.Request().Context()
	if err := h.ExpenseService.Delete(ctx, id); err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

// ListExpenses godoc
//
//	@Summary		List all expenses
//	@Description	get all expenses
//	@Tags			expenses
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		domain.Expense
//	@Failure		500	{object}	ResponseError
//	@Router			/expenses [get]
func (h *ExpenseHandler) ListExpenses(c echo.Context) error {
	ctx := c.Request().Context()
	expenses, err := h.ExpenseService.GetAll(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, expenses)
}

// GetExpensesByUserID godoc
//
//	@Summary		Get expenses by user ID
//	@Description	get expenses by user ID
//	@Tags			expenses
//	@Accept			json
//	@Produce		json
//	@Param			user_id	path		int	true	"User ID"
//	@Success		200		{array}		domain.Expense
//	@Failure		400		{object}	ResponseError
//	@Failure		404		{object}	ResponseError
//	@Router			/users/{user_id}/expenses [get]
func (h *ExpenseHandler) GetExpensesByUserID(c echo.Context) error {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: "invalid user id"})
	}

	ctx := c.Request().Context()
	expenses, err := h.ExpenseService.GetByUserID(ctx, userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, expenses)
}
