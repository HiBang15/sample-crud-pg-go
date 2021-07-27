package repository

import "github.com/HiBang15/sample-crud-pg-go.git/model"

type IPermissionRepository interface {
	GetListPermission() ([]*model.Permission, error)
	GetPermissionById(int) (*model.Permission, error)
	CreatePermission(*model.Permission) error
	UpdatePermission(*model.Permission) error
	DeletePermission(int) error
}
