package http

import (
	"context"
	"gamelieelearn/expense-tracker-api-go/domain"
	"net/http"
)

type ExpenseService interface {
	Store(ctx context.Context, expense *domain.Expense) (err error)
}

type ExpenseHandler struct {
	ExpenseService ExpenseService
}

func NewExpenseHandler(expenseService ExpenseService) ExpenseHandler {
	return ExpenseHandler{
		ExpenseService: expenseService,
	}
}

func (h *ExpenseHandler) GetExpense(w http.ResponseWriter, r *http.Request) {
	// Implement the handler
}
