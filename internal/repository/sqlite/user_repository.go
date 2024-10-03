package sqlite

import (
	"context"
	"database/sql"
	"gamelieelearn/expense-tracker-api-go/domain"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetByID(ctx context.Context, id int64) (domain.User, error) {
	return domain.User{}, nil
}

func (r *UserRepository) Store(ctx context.Context, u *domain.User) error {
	return nil
}

func (r *UserRepository) Hello() error {
	return nil
}
