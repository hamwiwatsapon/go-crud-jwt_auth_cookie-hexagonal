package core

// PORT FOR SECONDARY ADAPTERS
type UserRepository interface {
	CreateUser(user User) error
	ReadUser(user User) 	(User, error)
	UpdateUser(user User) error
	DeleteUser(user User) error
}