package adapters

import (
	core "github.com/hamwiwatsapon/go-crud-authen/core"
	"gorm.io/gorm"
)

func NewGormUser(db *gorm.DB) core.UserRepository {
	return &GormBookStore{db: db}
}

func (r *GormBookStore) CreateUser(user core.User) error {
	if result := r.db.Create(&user); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *GormBookStore) ReadUser(user core.User) (core.User, error) {
	selectedUser := new(core.User)
	result := r.db.Where(&core.User{Email: user.Email, Username: user.Username}).First(selectedUser)

	if result.Error != nil {
		return *selectedUser, result.Error
	}

	return *selectedUser, nil
}

func (r *GormBookStore) UpdateUser(user core.User) error {
	result := r.db.Model(&user).Where(&core.User{Email: user.Email, Username: user.Username}).Update("password", user.Password)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *GormBookStore) DeleteUser(user core.User) error {
	selectedUser := new(core.User)
	result := r.db.Where(&core.User{Email: user.Email, Username: user.Username}).Delete(selectedUser)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
