package adapters

import (
	"bytes"
	"errors"
	"testing"

	"net/http/httptest"

	"github.com/gofiber/fiber/v2"
	"github.com/hamwiwatsapon/go-crud-authen/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

// Authentication implements core.UserService.
func (m *MockUserService) Authentication(user core.User) (string, error) {
	args := m.Called(user)
	return "test_mock", args.Error(0)
}

// Delete implements core.UserService.
func (m *MockUserService) Delete(user core.User) error {
	args := m.Called(user)
	return args.Error(0)
}

// Edit implements core.UserService.
func (m *MockUserService) Edit(user core.User) error {
	args := m.Called(user)
	return args.Error(0)
}

// Register implements core.UserService.
func (m *MockUserService) Register(user core.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func TestRegisterUserHandler(t *testing.T) {
	mockService := new(MockUserService)
	handler := NewHttpUserHandler(mockService)

	app := fiber.New()
	registerUrl := "/register"

	app.Post(registerUrl, handler.Register)

	t.Run("successful user register", func(t *testing.T) {
		mockService.On("Register", mock.AnythingOfType("core.User")).Return(nil)

		req := httptest.NewRequest("POST", registerUrl, bytes.NewBufferString(`{"username": "test_test","email":"test@test.com","password":"test1234"}`))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("fail user register (password less than 8)", func(t *testing.T) {
		mockService.On("Register", mock.AnythingOfType("core.User")).Return(errors.New("password is equal or more than 8 characters"))

		req := httptest.NewRequest("POST", registerUrl, bytes.NewBufferString(`{"username": "test_test","email":"test@test.com","password":"test12"}`))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}
