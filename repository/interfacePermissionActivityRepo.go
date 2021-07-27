package repository

import "github.com/HiBang15/sample-crud-pg-go.git/model"

type IPermissionActivityRepository interface {
	GetListPermissionActivity() ([]*model.PermissionActivity, error)
	GetPermissionActivityByPermissionId(int) (*model.PermissionActivity, error)
	CreatePermissionActivity(*model.PermissionActivity) error
	UpdatePermissionActivityByPermissionId(*model.PermissionActivity) error
	DeletePermissionActivityByPermissionId(int) error

	GetActivityByPermissionId(int)([]*model.Permission, error)
}
