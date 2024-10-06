package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"gamelieelearn/expense-tracker-api-go/domain"
	"time"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetByID(ctx context.Context, id int64) (domain.User, error) {
	query := `SELECT id, name, created_at, updated_at FROM users WHERE id = ?`
	row := r.DB.QueryRowContext(ctx, query, id)

	var user domain.User
	err := row.Scan(&user.ID, &user.Name, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.User{}, errors.New("user not found")
		}
		return domain.User{}, err
	}

	return user, nil
}

func (r *UserRepository) Store(ctx context.Context, u *domain.User) error {
	query := `INSERT INTO users (name, created_at, updated_at) VALUES (?, ?, ?)`
	now := time.Now().Format(time.RFC3339)
	result, err := r.DB.ExecContext(ctx, query, u.Name, now, now)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	u.ID = id
	u.CreatedAt = now
	u.UpdatedAt = now
	return nil
}

func (r *UserRepository) Update(ctx context.Context, u *domain.User) error {
	query := `UPDATE users SET name = ?, updated_at = ? WHERE id = ?`
	now := time.Now().Format(time.RFC3339)
	_, err := r.DB.ExecContext(ctx, query, u.Name, now, u.ID)
	if err != nil {
		return err
	}

	u.UpdatedAt = now
	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM users WHERE id = ?`
	_, err := r.DB.ExecContext(ctx, query, id)
	return err
}

func (r *UserRepository) GetAll(ctx context.Context) ([]domain.User, error) {
	query := `SELECT id, name, created_at, updated_at FROM users`
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return []domain.User{}, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.ID, &user.Name, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return []domain.User{}, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return []domain.User{}, err
	}

	return users, nil
}
