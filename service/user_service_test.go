package service_test

import (
	"context"
	"gamelieelearn/expense-tracker-api-go/domain"
	"gamelieelearn/expense-tracker-api-go/service"
	"gamelieelearn/expense-tracker-api-go/service/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserGetById(t *testing.T) {
	mockUserRepository := mocks.NewUserRepository(t)
	mockUser := domain.User{
		ID:        1,
		Name:      "mockUser",
		CreatedAt: "2019-03-07 15:08:00 +0000 UTC",
		UpdatedAt: "2019-03-07 15:08:00 +0000 UTC",
	}
	t.Run("simple get", func(t *testing.T) {
		mockUserRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockUser, nil)
		userService := &service.UserService{UserRepository: mockUserRepository}

		user, err := userService.GetByID(context.TODO(), 1)
		assert.NoError(t, err)
		assert.NotNil(t, user)

		mockUserRepository.AssertExpectations(t)
	})
}

func TestUserStore(t *testing.T) {
	mockUserRepository := mocks.NewUserRepository(t)
	mockUser := domain.User{
		ID:        1,
		Name:      "mockUser",
		CreatedAt: "2019-03-07 15:08:00 +0000 UTC",
		UpdatedAt: "2019-03-07 15:08:00 +0000 UTC",
	}
	t.Run("success", func(t *testing.T) {
		mockUserRepository.On("Store", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil).Once()
		tempMockUser := mockUser
		tempMockUser.ID = 0

		userService := &service.UserService{UserRepository: mockUserRepository}

		err := userService.Store(context.TODO(), &mockUser)

		assert.NoError(t, err)
		assert.Equal(t, mockUser.Name, tempMockUser.Name)
	})
}

func TestUserUpdate(t *testing.T) {
	mockUserRepository := mocks.NewUserRepository(t)
	mockUser := domain.User{
		ID:        1,
		Name:      "updatedUser",
		CreatedAt: "2019-03-07 15:08:00 +0000 UTC",
		UpdatedAt: "2019-03-08 15:08:00 +0000 UTC",
	}
	t.Run("success", func(t *testing.T) {
		mockUserRepository.On("Update", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil).Once()

		userService := &service.UserService{UserRepository: mockUserRepository}

		err := userService.Update(context.TODO(), &mockUser)

		assert.NoError(t, err)
		mockUserRepository.AssertExpectations(t)
	})
}

func TestUserDelete(t *testing.T) {
	mockUserRepository := mocks.NewUserRepository(t)
	t.Run("success", func(t *testing.T) {
		mockUserRepository.On("Delete", mock.Anything, mock.AnythingOfType("int64")).Return(nil).Once()

		userService := &service.UserService{UserRepository: mockUserRepository}

		err := userService.Delete(context.TODO(), 1)

		assert.NoError(t, err)
		mockUserRepository.AssertExpectations(t)
	})
}

func TestUserGetAll(t *testing.T) {
	mockUserRepository := mocks.NewUserRepository(t)
	mockUsers := []domain.User{
		{
			ID:        1,
			Name:      "user1",
			CreatedAt: "2019-03-07 15:08:00 +0000 UTC",
			UpdatedAt: "2019-03-07 15:08:00 +0000 UTC",
		},
		{
			ID:        2,
			Name:      "user2",
			CreatedAt: "2019-03-08 15:08:00 +0000 UTC",
			UpdatedAt: "2019-03-08 15:08:00 +0000 UTC",
		},
	}
	t.Run("success", func(t *testing.T) {
		mockUserRepository.On("GetAll", mock.Anything).Return(mockUsers, nil).Once()

		userService := &service.UserService{UserRepository: mockUserRepository}

		users, err := userService.GetAll(context.TODO())

		assert.NoError(t, err)
		assert.Len(t, users, 2)
		assert.Equal(t, mockUsers, users)
		mockUserRepository.AssertExpectations(t)
	})
}
