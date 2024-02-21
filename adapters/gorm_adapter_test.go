package adapters

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/hamwiwatsapon/go-crud-authen/core"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGormRepository(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	repo := NewGormUser(gormDB)

	t.Run("success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO  "users"`).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()
		err := repo.CreateUser(core.User{
			Username: "test_user",
			Email:    "test_usertest.com",
			Password: "test1234",
		})
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("failure", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO  "users"`).WillReturnError(errors.New("database error"))
		mock.ExpectRollback()
		err := repo.CreateUser(core.User{
			Username: "test_user",
			Email:    "test_usertest.com",
			Password: "test1234",
		})
		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
