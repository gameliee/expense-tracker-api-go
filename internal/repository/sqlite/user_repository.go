package sqlite

import (
	"context"
	"errors"
	"gamelieelearn/expense-tracker-api-go/domain"
	"time"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB `inject:"*gorm.DB"`
}

func (r *UserRepository) GetByID(ctx context.Context, id int64) (user domain.User, err error) {
	r.DB.First(&user, id)
	if user == (domain.User{}) {
		err = errors.New("user not found")
	}

	return user, err
}

func (r *UserRepository) Store(ctx context.Context, u *domain.User) error {
	now := time.Now().Format(time.RFC3339)
	u.CreatedAt = now
	u.UpdatedAt = now
	result := r.DB.Create(u)
	return result.Error
}

func (r *UserRepository) Update(ctx context.Context, u *domain.User) error {
	result := r.DB.First(u)
	if result.Error != nil {
		return result.Error
	}
	u.UpdatedAt = time.Now().Format(time.RFC3339)
	result = r.DB.Save(u)
	return result.Error
}

func (r *UserRepository) Delete(ctx context.Context, id int64) error {
	result := r.DB.Delete(&domain.User{}, id)
	return result.Error
}

func (r *UserRepository) GetAll(ctx context.Context) (users []domain.User, err error) {
	result := r.DB.Find(&users)
	err = result.Error
	return
}
