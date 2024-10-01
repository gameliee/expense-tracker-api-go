package sqlite

import (
	"context"
	"gamelieelearn/expense-tracker-api-go/domain"
)

type UserRepository struct {
	// Add any necessary fields
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
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
