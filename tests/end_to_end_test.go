package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gamelieelearn/expense-tracker-api-go/cmd"
	"gamelieelearn/expense-tracker-api-go/domain"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var testServer *httptest.Server

func TestMain(m *testing.M) {
	// Set up the test environment
	os.Setenv("APP_MODE", "test")
	os.Setenv("DATABASE_PATH", ":memory:")

	cmd.InitializeApplication()
	container := cmd.GetContainer()

	testEcho := container.Get((*echo.Echo)(nil)).(*echo.Echo)

	// Start test server
	testServer = httptest.NewServer(testEcho.Server.Handler)
	defer testServer.Close()

	// Run the tests
	code := m.Run()

	// Clean up
	os.Exit(code)
}

func TestEndToEnd(t *testing.T) {
	// Test user creation
	user := domain.User{
		Name: "John Doe",
	}
	userJSON, _ := json.Marshal(user)
	resp, body := makeRequest(t, "POST", "/users", userJSON)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var createdUser domain.User
	err := json.Unmarshal(body, &createdUser)
	assert.NoError(t, err)
	assert.NotEmpty(t, createdUser.ID)

	// Test get user
	resp, body = makeRequest(t, "GET", fmt.Sprintf("/users/%d", createdUser.ID), nil)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	var fetchedUser domain.User
	err = json.Unmarshal(body, &fetchedUser)
	assert.NoError(t, err)
	assert.Equal(t, createdUser.ID, fetchedUser.ID)

	// Test update user
	updatedUser := domain.User{
		Name: "Jane Doe",
	}
	updatedUserJSON, _ := json.Marshal(updatedUser)
	resp, body = makeRequest(t, "PUT", fmt.Sprintf("/users/%d", createdUser.ID), updatedUserJSON)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	err = json.Unmarshal(body, &fetchedUser)
	assert.NoError(t, err)
	assert.Equal(t, "Jane Doe", fetchedUser.Name)

	// Test list users
	resp, body = makeRequest(t, "GET", "/users", nil)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	var users []domain.User
	err = json.Unmarshal(body, &users)
	assert.NoError(t, err)
	assert.Len(t, users, 1)

	// Test create expense
	expense := domain.Expense{
		User_ID:     createdUser.ID,
		Amount:      100.0,
		Name:        "Test Expense",
		Description: "Test expense",
	}
	expenseJSON, _ := json.Marshal(expense)
	resp, body = makeRequest(t, "POST", "/expenses", expenseJSON)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var createdExpense domain.Expense
	err = json.Unmarshal(body, &createdExpense)
	assert.NoError(t, err)
	assert.NotEmpty(t, createdExpense.ID)

	// Test get expense
	resp, body = makeRequest(t, "GET", fmt.Sprintf("/expenses/%d", createdExpense.ID), nil)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	var fetchedExpense domain.Expense
	err = json.Unmarshal(body, &fetchedExpense)
	assert.NoError(t, err)
	assert.Equal(t, createdExpense.ID, fetchedExpense.ID)

	// Test update expense
	updatedExpense := domain.Expense{
		User_ID:     createdUser.ID,
		Amount:      150.0,
		Name:        "Updated Test Expense",
		Description: "Updated test expense",
	}
	updatedExpenseJSON, _ := json.Marshal(updatedExpense)
	resp, body = makeRequest(t, "PUT", fmt.Sprintf("/expenses/%d", createdExpense.ID), updatedExpenseJSON)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	err = json.Unmarshal(body, &fetchedExpense)
	assert.NoError(t, err)
	assert.Equal(t, 150.0, fetchedExpense.Amount)
	assert.Equal(t, "Updated test expense", fetchedExpense.Description)

	// Test list expenses
	resp, body = makeRequest(t, "GET", "/expenses", nil)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	var expenses []domain.Expense
	err = json.Unmarshal(body, &expenses)
	assert.NoError(t, err)
	assert.Len(t, expenses, 1)

	// Test list expenses for user
	resp, body = makeRequest(t, "GET", fmt.Sprintf("/users/%d/expenses", createdUser.ID), nil)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	err = json.Unmarshal(body, &expenses)
	assert.NoError(t, err)
	assert.Len(t, expenses, 1)
	assert.Equal(t, createdExpense.ID, expenses[0].ID)

	// Test delete expense
	resp, _ = makeRequest(t, "DELETE", fmt.Sprintf("/expenses/%d", createdExpense.ID), nil)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)

	// Verify expense is deleted
	resp, _ = makeRequest(t, "GET", fmt.Sprintf("/expenses/%d", createdExpense.ID), nil)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)

	// Test delete user
	resp, _ = makeRequest(t, "DELETE", fmt.Sprintf("/users/%d", createdUser.ID), nil)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)

	// Verify user is deleted
	resp, _ = makeRequest(t, "GET", fmt.Sprintf("/users/%d", createdUser.ID), nil)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func makeRequest(t *testing.T, method, path string, body []byte) (*http.Response, []byte) {
	req, err := http.NewRequest(method, testServer.URL+path, bytes.NewBuffer(body))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(t, err)

	respBody, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	defer resp.Body.Close()

	return resp, respBody
}
