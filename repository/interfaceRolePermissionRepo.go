package repository

import "github.com/HiBang15/sample-crud-pg-go.git/model"

type IRolePermissionRepository interface {
	GetListRolePermission() ([]*model.RolePermission, error)
	GetRolePermissionByRoleId(int) (*model.RolePermission, error)
	GetRolePermissionByPermissionId(int) (*model.RolePermission, error)
	CreateRolePermission(*model.RolePermission) error
	UpdateRolePermissionByRoleId(*model.RolePermission) error
	DeleteRolePermissionByPermissionId(int) error

	GetPermissionByRoleId(int)([]*model.Permission, error)
}
