package sqlite

import (
	"context"
	"errors"
	"gamelieelearn/expense-tracker-api-go/domain"
	"time"

	"gorm.io/gorm"
)

type ExpenseRepository struct {
	DB *gorm.DB
}

func NewExpenseRepository(db *gorm.DB) *ExpenseRepository {
	return &ExpenseRepository{DB: db}
}

func (r *ExpenseRepository) Store(ctx context.Context, expense *domain.Expense) error {
	now := time.Now().Format(time.RFC3339)
	expense.CreatedAt = now
	expense.UpdatedAt = now
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
	expense.UpdatedAt = time.Now().Format(time.RFC3339)
	result := r.DB.Save(expense)
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
