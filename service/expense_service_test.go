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
	mockExpenseRepo := new(mocks.ExpenseRepository)
	mockUserRepo := new(mocks.UserRepository)
	userService := &service.UserService{UserRepository: mockUserRepo}
	expenseService := &service.ExpenseService{ExpenseRepository: mockExpenseRepo, UserService: userService}

	mockExpense := &domain.Expense{
		ID:          1,
		User_ID:     1,
		Amount:      100.0,
		Description: "Test expense",
		CreatedAt:   "2023-05-01 10:00:00",
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("GetByID", mock.Anything, int64(1)).Return(domain.User{ID: 1}, nil).Once()
		mockExpenseRepo.On("Store", mock.Anything, mockExpense).Return(nil)

		err := expenseService.Store(context.TODO(), mockExpense)

		assert.NoError(t, err)
		mockUserRepo.AssertExpectations(t)
		mockExpenseRepo.AssertExpectations(t)
	})

	t.Run("user not found", func(t *testing.T) {
		mockUserRepo.On("GetByID", mock.Anything, int64(1)).Return(domain.User{}, errors.New("user not found")).Once()

		err := expenseService.Store(context.TODO(), mockExpense)

		assert.Error(t, err)
		assert.EqualError(t, err, "user not found")
		mockUserRepo.AssertExpectations(t)
		mockExpenseRepo.AssertNotCalled(t, "Store")
	})
}

func TestExpenseGetByID(t *testing.T) {
	mockExpenseRepo := new(mocks.ExpenseRepository)
	mockUserRepo := new(mocks.UserRepository)
	userService := &service.UserService{UserRepository: mockUserRepo}
	expenseService := &service.ExpenseService{ExpenseRepository: mockExpenseRepo, UserService: userService}

	mockExpense := domain.Expense{
		ID:          1,
		User_ID:     1,
		Amount:      100.0,
		Description: "Test expense",
		CreatedAt:   "2023-05-01 10:00:00",
	}

	t.Run("success", func(t *testing.T) {
		mockExpenseRepo.On("GetByID", mock.Anything, int64(1)).Return(mockExpense, nil)

		expense, err := expenseService.GetByID(context.TODO(), 1)

		assert.NoError(t, err)
		assert.Equal(t, mockExpense, expense)
		mockExpenseRepo.AssertExpectations(t)
	})
}

func TestExpenseUpdate(t *testing.T) {
	mockExpenseRepo := new(mocks.ExpenseRepository)
	mockUserRepo := new(mocks.UserRepository)
	userService := &service.UserService{UserRepository: mockUserRepo}
	expenseService := &service.ExpenseService{ExpenseRepository: mockExpenseRepo, UserService: userService}

	mockExpense := &domain.Expense{
		ID:          1,
		User_ID:     1,
		Amount:      150.0,
		Description: "Updated test expense",
		CreatedAt:   "2023-05-01 10:00:00",
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("GetByID", mock.Anything, int64(1)).Return(domain.User{ID: 1}, nil).Once()
		mockExpenseRepo.On("Update", mock.Anything, mockExpense).Return(nil)

		err := expenseService.Update(context.TODO(), mockExpense)

		assert.NoError(t, err)
		mockUserRepo.AssertExpectations(t)
		mockExpenseRepo.AssertExpectations(t)
	})

	t.Run("user not found", func(t *testing.T) {
		mockUserRepo.On("GetByID", mock.Anything, int64(1)).Return(domain.User{}, errors.New("user not found")).Once()

		err := expenseService.Update(context.TODO(), mockExpense)

		assert.Error(t, err)
		assert.EqualError(t, err, "user not found")
		mockUserRepo.AssertExpectations(t)
		mockExpenseRepo.AssertNotCalled(t, "Update")
	})
}

func TestExpenseDelete(t *testing.T) {
	mockExpenseRepo := new(mocks.ExpenseRepository)
	mockUserRepo := new(mocks.UserRepository)
	userService := &service.UserService{UserRepository: mockUserRepo}
	expenseService := &service.ExpenseService{ExpenseRepository: mockExpenseRepo, UserService: userService}

	t.Run("success", func(t *testing.T) {
		mockExpenseRepo.On("Delete", mock.Anything, int64(1)).Return(nil)

		err := expenseService.Delete(context.TODO(), 1)

		assert.NoError(t, err)
		mockExpenseRepo.AssertExpectations(t)
	})
}

func TestExpenseGetAll(t *testing.T) {
	mockExpenseRepo := new(mocks.ExpenseRepository)
	mockUserRepo := new(mocks.UserRepository)
	userService := &service.UserService{UserRepository: mockUserRepo}
	expenseService := &service.ExpenseService{ExpenseRepository: mockExpenseRepo, UserService: userService}

	mockExpenses := []domain.Expense{
		{ID: 1, User_ID: 1, Amount: 100.0, Description: "Expense 1", CreatedAt: "2023-05-01 10:00:00"},
		{ID: 2, User_ID: 1, Amount: 200.0, Description: "Expense 2", CreatedAt: "2023-05-02 10:00:00"},
	}

	t.Run("success", func(t *testing.T) {
		mockExpenseRepo.On("GetAll", mock.Anything).Return(mockExpenses, nil)

		expenses, err := expenseService.GetAll(context.TODO())

		assert.NoError(t, err)
		assert.Equal(t, mockExpenses, expenses)
		mockExpenseRepo.AssertExpectations(t)
	})
}

func TestExpenseGetByUserID(t *testing.T) {
	mockExpenseRepo := new(mocks.ExpenseRepository)
	mockUserRepo := new(mocks.UserRepository)
	userService := &service.UserService{UserRepository: mockUserRepo}
	expenseService := &service.ExpenseService{ExpenseRepository: mockExpenseRepo, UserService: userService}

	mockExpenses := []domain.Expense{
		{ID: 1, User_ID: 1, Amount: 100.0, Description: "Expense 1", CreatedAt: "2023-05-01 10:00:00"},
		{ID: 2, User_ID: 1, Amount: 200.0, Description: "Expense 2", CreatedAt: "2023-05-02 10:00:00"},
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("GetByID", mock.Anything, int64(1)).Return(domain.User{ID: 1}, nil).Once()
		mockExpenseRepo.On("GetByUserID", mock.Anything, int64(1)).Return(mockExpenses, nil)

		expenses, err := expenseService.GetByUserID(context.TODO(), 1)

		assert.NoError(t, err)
		assert.Equal(t, mockExpenses, expenses)
		mockUserRepo.AssertExpectations(t)
		mockExpenseRepo.AssertExpectations(t)
	})

	t.Run("user not found", func(t *testing.T) {
		mockUserRepo.On("GetByID", mock.Anything, int64(1)).Return(domain.User{}, errors.New("user not found")).Once()

		expenses, err := expenseService.GetByUserID(context.TODO(), 1)

		assert.Error(t, err)
		assert.EqualError(t, err, "user not found")
		assert.Nil(t, expenses)
		mockUserRepo.AssertExpectations(t)
		mockExpenseRepo.AssertNotCalled(t, "GetByUserID")
	})
}
