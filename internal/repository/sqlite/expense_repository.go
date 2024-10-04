package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"gamelieelearn/expense-tracker-api-go/domain"
	"time"
)

type ExpenseRepository struct {
	DB *sql.DB
}

func NewExpenseRepository(db *sql.DB) *ExpenseRepository {
	return &ExpenseRepository{DB: db}
}

func (r *ExpenseRepository) Store(ctx context.Context, expense *domain.Expense) error {
	query := `INSERT INTO expenses (user_id, name, description, amount, created_at) VALUES (?, ?, ?, ?, ?)`
	now := time.Now().Format(time.RFC3339)
	result, err := r.DB.ExecContext(ctx, query, expense.User_ID, expense.Name, expense.Description, expense.Amount, now)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	expense.ID = id
	expense.CreatedAt = now
	return nil
}

func (r *ExpenseRepository) GetByID(ctx context.Context, id int64) (domain.Expense, error) {
	query := `SELECT id, user_id, name, description, amount, created_at FROM expenses WHERE id = ?`
	row := r.DB.QueryRowContext(ctx, query, id)

	var expense domain.Expense
	err := row.Scan(&expense.ID, &expense.User_ID, &expense.Name, &expense.Description, &expense.Amount, &expense.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Expense{}, errors.New("expense not found")
		}
		return domain.Expense{}, err
	}

	return expense, nil
}

func (r *ExpenseRepository) Update(ctx context.Context, expense *domain.Expense) error {
	query := `UPDATE expenses SET user_id = ?, name = ?, description = ?, amount = ? WHERE id = ?`
	_, err := r.DB.ExecContext(ctx, query, expense.User_ID, expense.Name, expense.Description, expense.Amount, expense.ID)
	return err
}

func (r *ExpenseRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM expenses WHERE id = ?`
	_, err := r.DB.ExecContext(ctx, query, id)
	return err
}

func (r *ExpenseRepository) GetAll(ctx context.Context) ([]domain.Expense, error) {
	query := `SELECT id, user_id, name, description, amount, created_at FROM expenses`
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []domain.Expense
	for rows.Next() {
		var expense domain.Expense
		err := rows.Scan(&expense.ID, &expense.User_ID, &expense.Name, &expense.Description, &expense.Amount, &expense.CreatedAt)
		if err != nil {
			return nil, err
		}
		expenses = append(expenses, expense)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return expenses, nil
}

func (r *ExpenseRepository) GetByUserID(ctx context.Context, userID int64) ([]domain.Expense, error) {
	query := `SELECT id, user_id, name, description, amount, created_at FROM expenses WHERE user_id = ?`
	rows, err := r.DB.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []domain.Expense
	for rows.Next() {
		var expense domain.Expense
		err := rows.Scan(&expense.ID, &expense.User_ID, &expense.Name, &expense.Description, &expense.Amount, &expense.CreatedAt)
		if err != nil {
			return nil, err
		}
		expenses = append(expenses, expense)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return expenses, nil
}
