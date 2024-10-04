package service

import (
	"context"
	"gamelieelearn/expense-tracker-api-go/domain"
)

//go:generate mockery --name UserRepositoryInt --structname UserRepository
type UserRepositoryInt interface {
	Store(ctx context.Context, u *domain.User) error
	GetByID(ctx context.Context, id int64) (domain.User, error)
	Update(ctx context.Context, u *domain.User) error
	Delete(ctx context.Context, id int64) error
	GetAll(ctx context.Context) ([]domain.User, error)
}

type UserService struct {
	UserRepository UserRepositoryInt
}

func NewUserService(userRepository UserRepositoryInt) *UserService {
	return &UserService{UserRepository: userRepository}
}

func (s *UserService) Store(ctx context.Context, u *domain.User) error {
	return s.UserRepository.Store(ctx, u)
}

func (s *UserService) GetByID(ctx context.Context, id int64) (domain.User, error) {
	return s.UserRepository.GetByID(ctx, id)
}

func (s *UserService) Update(ctx context.Context, u *domain.User) error {
	return s.UserRepository.Update(ctx, u)
}

func (s *UserService) Delete(ctx context.Context, id int64) error {
	return s.UserRepository.Delete(ctx, id)
}

func (s *UserService) GetAll(ctx context.Context) ([]domain.User, error) {
	return s.UserRepository.GetAll(ctx)
}
