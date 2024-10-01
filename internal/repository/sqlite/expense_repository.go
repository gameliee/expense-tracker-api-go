package sqlite

import (
	"context"
	"gamelieelearn/expense-tracker-api-go/domain"
)

type ExpenseRepository struct {
	// Add any necessary fields
}

func NewExpenseRepository() *ExpenseRepository {
	return &ExpenseRepository{}
}

func (r *ExpenseRepository) Store(ctx context.Context, expense *domain.Expense) error {
	return nil
}
