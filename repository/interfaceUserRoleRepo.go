package repository

import "github.com/HiBang15/sample-crud-pg-go.git/model"

type IUserRoleRepository interface {
	GetListUserRole() ([]*model.UserRole, error)
	GetUserRoleByRoleId(int) (*model.UserRole, error)
	GetUserRoleByUserId(int) (*model.UserRole, error)
	CreateUserRole(*model.UserRole) error
	UpdateUserRoleByUserId(*model.UserRole) error
	UpdateUserRoleByRoleId(*model.UserRole) error
	DeleteUserRoleByUserId(int) error
	DeleteUserRoleByRoleId(int) error

	GetRoleByUserId(int) ([]*model.Role, error)
}
