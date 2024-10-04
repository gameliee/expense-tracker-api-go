package service_test

import (
	"context"
	"errors"
	"gamelieelearn/expense-tracker-api-go/domain"
	"gamelieelearn/expense-tracker-api-go/service"
	"gamelieelearn/expense-tracker-api-go/service/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestExpenseStore(t *testing.T) {
	mockExpenseRepository := mocks.NewExpenseRepository(t)
	mockExpenseRepository.On("Store", mock.Anything, mock.AnythingOfType("*domain.Expense")).Return(nil).Once()
	mockExpense := domain.Expense{
		ID:          100,
		User_ID:     1,
		Amount:      1234,
		Description: "mockDescription",
		CreatedAt:   "2019-03-07T15:08:00+00:00",
	}
	t.Run("success", func(t *testing.T) {
		mockUserRepository := mocks.NewUserRepository(t)
		mockUser := domain.User{
			ID:        1,
			Name:      "mockUser",
			CreatedAt: "2019-03-07 15:08:00 +0000 UTC",
			UpdatedAt: "2019-03-07 15:08:00 +0000 UTC",
		}
		mockUserRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockUser, nil)

		userService := service.NewUserService(mockUserRepository)
		expenseService, err := service.NewExpenseService(mockExpenseRepository, userService)
		assert.NoError(t, err)

		err = expenseService.Store(context.TODO(), &mockExpense)

		assert.NoError(t, err)
	})
	t.Run("user do not exist", func(t *testing.T) {
		mockUserRepository := mocks.NewUserRepository(t)
		mockUserRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(domain.User{}, errors.New("user not found"))

		userService := service.NewUserService(mockUserRepository)
		expenseService, err := service.NewExpenseService(mockExpenseRepository, userService)
		assert.NoError(t, err)

		err = expenseService.Store(context.TODO(), &mockExpense)
		assert.Error(t, err)
	})
}
