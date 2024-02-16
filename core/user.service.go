package core

import (
	"errors"
	"net/mail"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// PORT FOR PRIMARY ADAPTERS
type UserService interface {
	Register(user User) error
	Authentication(user User) (string, error)
	Edit(user User) error
	Delete(user User) error
}

type userServiceImpl struct {
	repo UserRepository
}

func (s *userServiceImpl) Register(user User) error {
	if len(user.Password) < 8 {
		return errors.New("password is equal or more than 8 characters")
	}

	if _, err := mail.ParseAddress(user.Email); err != nil {
		return err
	}

	if user.Username == "" || len(user.Username) < 8 {
		return errors.New("username cant be empty and equal or more than 8 characters")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	if err := s.repo.CreateUser(user); err != nil {
		return err
	}

	return nil
}

func (s *userServiceImpl) Authentication(user User) (string, error) {
	selectedUser, err := s.repo.ReadUser(user)

	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(selectedUser.Password),
		[]byte(user.Password),
	)

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": selectedUser.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	jwtSecretKey := os.Getenv("jwtSeCretKey")
	tokenString, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *userServiceImpl) Edit(user User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	if err := s.repo.UpdateUser(user); err != nil {
		return err
	}
	return nil
}

func (s *userServiceImpl) Delete(user User) error {
	if err := s.repo.DeleteUser(user); err != nil {
		return err
	}
	return nil
}

func NewUserService(repo UserRepository) UserService {
	return &userServiceImpl{repo: repo}
}
