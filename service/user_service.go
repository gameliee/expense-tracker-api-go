package service

import (
	"context"
	"gamelieelearn/expense-tracker-api-go/domain"
)

type UserRepository interface {
	Store(ctx context.Context, u *domain.User) error
}

type UserService struct {
	userRepository UserRepository
}

func (s *UserService) Store(ctx context.Context, u *domain.User) (err error) {
	err = s.userRepository.Store(ctx, u)
	return
}
