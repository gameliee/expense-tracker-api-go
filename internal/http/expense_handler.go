package http

import (
	"context"
	"gamelieelearn/expense-tracker-api-go/domain"
	"net/http"
)

type ExpenseServiceInt interface {
	Store(ctx context.Context, expense *domain.Expense) (err error)
}

type ExpenseHandler struct {
	ExpenseService ExpenseServiceInt
}

func NewExpenseHandler(expenseService ExpenseServiceInt) *ExpenseHandler {
	return &ExpenseHandler{
		ExpenseService: expenseService,
	}
}

func (h *ExpenseHandler) GetExpense(w http.ResponseWriter, r *http.Request) {
	// Implement the handler
}
