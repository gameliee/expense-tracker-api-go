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

func TestGetById(t *testing.T) {

	mockUserRepository := mocks.NewUserRepository(t)
	mockUser := domain.User{
		ID:        1,
		Name:      "mockUser",
		CreatedAt: "2019-03-07 15:08:00 +0000 UTC",
		UpdatedAt: "2019-03-07 15:08:00 +0000 UTC",
	}
	t.Run("simple get", func(t *testing.T) {
		mockUserRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockUser, nil)

		userService := service.NewUserService(mockUserRepository)
		user, err := userService.GetByID(context.TODO(), 1)
		assert.NoError(t, err)
		assert.NotNil(t, user)

		mockUserRepository.AssertExpectations(t)
	})
}

func TestStore(t *testing.T) {
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

		userService := service.NewUserService(mockUserRepository)
		err := userService.Store(context.TODO(), &mockUser)

		assert.NoError(t, err)
		assert.Equal(t, mockUser.Name, tempMockUser.Name)
	})
}
