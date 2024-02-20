package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockUserRepo struct {
	createUserFunc func(user User) error
	readUserFunc   func(user User) (User, error)
	updateUserFunc func(user User) error
	deleteUserFunc func(user User) error
}

func (m *mockUserRepo) DeleteUser(user User) error {
	return m.deleteUserFunc(user)
}

func (m *mockUserRepo) ReadUser(user User) (User, error) {
	return m.readUserFunc(user)
}

func (m *mockUserRepo) UpdateUser(user User) error {
	return m.updateUserFunc(user)
}

func (m *mockUserRepo) CreateUser(user User) error {
	return m.createUserFunc(user)
}

func TestCreateUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repo := &mockUserRepo{
			createUserFunc: func(user User) error {
				return nil
			},
		}

		service := NewUserService(repo)

		err := service.Register(User{
			Username: "test_user",
			Email:    "test_user_success@test.com",
			Password: "test_password",
		})
		assert.NoError(t, err)
	})

	t.Run("error password length", func(t *testing.T) {
		repo := &mockUserRepo{
			createUserFunc: func(user User) error {
				return nil
			},
		}

		service := NewUserService(repo)

		err := service.Register(User{
			Username: "test_user",
			Email:    "test_user_password@test.com",
			Password: "test",
		})

		assert.Error(t, err)
		assert.Equal(t, "password is equal or more than 8 characters", err.Error())
	})

	t.Run("error email missing @", func(t *testing.T) {
		repo := &mockUserRepo{
			createUserFunc: func(user User) error {
				return nil
			},
		}

		service := NewUserService(repo)

		err := service.Register(User{
			Username: "test_user",
			Email:    "test_usertest.com",
			Password: "test1234",
		})

		assert.Error(t, err)
		assert.Equal(t, "mail: missing '@' or angle-addr", err.Error())
	})

	t.Run("error user empty", func(t *testing.T) {
		repo := &mockUserRepo{
			createUserFunc: func(user User) error {
				return nil
			},
		}

		service := NewUserService(repo)

		err := service.Register(User{
			Username: "",
			Email:    "test_user_username@test.com",
			Password: "test1234",
		})

		assert.Error(t, err)
		assert.Equal(t, "username cant be empty and equal or more than 8 characters", err.Error())
	})
}
