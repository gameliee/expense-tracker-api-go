package service

import (
	"context"
	"gamelieelearn/expense-tracker-api-go/domain"
)

//go:generate mockery --name ExpenseRepositoryInt --structname ExpenseRepository
type ExpenseRepositoryInt interface {
	Store(ctx context.Context, expense *domain.Expense) error
	GetByID(ctx context.Context, id int64) (domain.Expense, error)
	Update(ctx context.Context, expense *domain.Expense) error
	Delete(ctx context.Context, id int64) error
	GetAll(ctx context.Context) ([]domain.Expense, error)
	GetByUserID(ctx context.Context, userID int64) ([]domain.Expense, error)
}

type ExpenseService struct {
	ExpenseRepository ExpenseRepositoryInt `inject:"*sqlite.ExpenseRepository"`
	UserService       *UserService         `inject:"*service.UserService"`
}

func (s *ExpenseService) Store(ctx context.Context, expense *domain.Expense) error {
	_, err := s.UserService.GetByID(ctx, expense.User_ID)
	if err != nil {
		return err
	}
	return s.ExpenseRepository.Store(ctx, expense)
}

func (s *ExpenseService) GetByID(ctx context.Context, id int64) (domain.Expense, error) {
	return s.ExpenseRepository.GetByID(ctx, id)
}

func (s *ExpenseService) Update(ctx context.Context, expense *domain.Expense) error {
	_, err := s.UserService.GetByID(ctx, expense.User_ID)
	if err != nil {
		return err
	}
	return s.ExpenseRepository.Update(ctx, expense)
}

func (s *ExpenseService) Delete(ctx context.Context, id int64) error {
	return s.ExpenseRepository.Delete(ctx, id)
}

func (s *ExpenseService) GetAll(ctx context.Context) ([]domain.Expense, error) {
	return s.ExpenseRepository.GetAll(ctx)
}

func (s *ExpenseService) GetByUserID(ctx context.Context, userID int64) ([]domain.Expense, error) {
	_, err := s.UserService.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return s.ExpenseRepository.GetByUserID(ctx, userID)
}
