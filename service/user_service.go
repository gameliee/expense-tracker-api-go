package service

import (
	"context"
	"gamelieelearn/expense-tracker-api-go/domain"
)

type UserRepositoryInt interface {
	Store(ctx context.Context, u *domain.User) error
	GetByID(ctx context.Context, id int64) (domain.User, error)
}

type UserService struct {
	UserRepository UserRepositoryInt
}

func NewUserService(userRepository UserRepositoryInt) *UserService {
	return &UserService{UserRepository: userRepository}
}

func (s *UserService) Store(ctx context.Context, u *domain.User) (err error) {
	err = s.UserRepository.Store(ctx, u)
	return
}

func (s *UserService) GetByID(ctx context.Context, id int64) (res domain.User, err error) {
	res, err = s.UserRepository.GetByID(ctx, id)
	return
}
