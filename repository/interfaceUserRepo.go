package repository

import (
	"github.com/HiBang15/sample-crud-pg-go.git/model"
)

type IUserRepository interface {
	GetUserByEmail(string) (*model.User, error)
	GetListUser() ([]*model.User, error)
	GetUserById(int) (*model.User, error)
	CreateUser(*model.User) error
	UpdateUser(*model.User) error
	DeleteUser(int) error
}