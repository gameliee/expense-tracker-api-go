package service

import (
	"context"
	"gamelieelearn/expense-tracker-api-go/domain"
)

//go:generate mockery --name ExpenseRepositoryInt --structname ExpenseRepository
type ExpenseRepositoryInt interface {
	Store(ctx context.Context, expense *domain.Expense) error
}

type ExpenseService struct {
	ExpenseRepository ExpenseRepositoryInt
	UserService       *UserService
}

func NewExpenseService(expenseRepository ExpenseRepositoryInt, userService *UserService) (*ExpenseService, error) {
	return &ExpenseService{ExpenseRepository: expenseRepository, UserService: userService}, nil
}

func (s *ExpenseService) Store(ctx context.Context, expense *domain.Expense) (err error) {
	_, err = s.UserService.GetByID(ctx, expense.User_ID)
	if err != nil {
		return
	}
	err = s.ExpenseRepository.Store(ctx, expense)
	return
}
