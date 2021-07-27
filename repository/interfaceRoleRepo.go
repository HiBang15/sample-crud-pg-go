package repository

import "github.com/HiBang15/sample-crud-pg-go.git/model"

type IRoleRepository interface {
	GetListRole() ([]*model.Role, error)
	GetRoleById(int) (*model.Role, error)
	CreateRole(*model.Role) error
	UpdateRole(*model.Role) error
	DeleteRole(int) error
}
