package sqlite

import (
	"context"
	"errors"
	"gamelieelearn/expense-tracker-api-go/domain"

	"gorm.io/gorm"
)

type ExpenseRepository struct {
	DB *gorm.DB `inject:"*gorm.DB"`
}

func (r *ExpenseRepository) Store(ctx context.Context, expense *domain.Expense) error {
	result := r.DB.Create(expense)
	return result.Error
}

func (r *ExpenseRepository) GetByID(ctx context.Context, id int64) (expense domain.Expense, err error) {
	result := r.DB.First(&expense, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return domain.Expense{}, errors.New("expense not found")
	}
	return expense, result.Error
}

func (r *ExpenseRepository) Update(ctx context.Context, expense *domain.Expense) error {
	existingExpense := domain.Expense{}
	result := r.DB.First(&existingExpense, expense.ID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New("expense not found")
	}

	updates := domain.Expense{
		User_ID:     expense.User_ID,
		Name:        expense.Name,
		Description: expense.Description,
		Amount:      expense.Amount,
	}

	result = r.DB.Model(&existingExpense).Updates(updates)
	return result.Error
}

func (r *ExpenseRepository) Delete(ctx context.Context, id int64) error {
	result := r.DB.Delete(&domain.Expense{}, id)
	return result.Error
}

func (r *ExpenseRepository) GetAll(ctx context.Context) (expenses []domain.Expense, err error) {
	result := r.DB.Find(&expenses)
	return expenses, result.Error
}

func (r *ExpenseRepository) GetByUserID(ctx context.Context, userID int64) ([]domain.Expense, error) {
	var expenses []domain.Expense
	result := r.DB.Where("user_id = ?", userID).Find(&expenses)
	return expenses, result.Error
}
