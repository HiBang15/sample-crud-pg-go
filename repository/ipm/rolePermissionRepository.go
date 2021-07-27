package ipm

import (
	"github.com/HiBang15/sample-crud-pg-go.git/model"
	"github.com/HiBang15/sample-crud-pg-go.git/repository"
	"github.com/go-pg/pg"
)

type RolePermissionRepository struct {
	Connector *pg.DB
}

func NewRolePermissionRepository(db *pg.DB) repository.IRolePermissionRepository {
	return &RolePermissionRepository{Connector: db}
}

func (r *RolePermissionRepository) GetListRolePermission() ([]*model.RolePermission, error) {
	panic("implement me")
}

func (r *RolePermissionRepository) GetRolePermissionByRoleId(roleId int) (*model.RolePermission, error) {
	panic("implement me")
}

func (r *RolePermissionRepository) GetRolePermissionByPermissionId(permissionId int) (*model.RolePermission, error) {
	panic("implement me")
}

func (r *RolePermissionRepository) CreateRolePermission(permission *model.RolePermission) error {
	panic("implement me")
}

func (r *RolePermissionRepository) UpdateRolePermissionByRoleId(permission *model.RolePermission) error {
	panic("implement me")
}

func (r *RolePermissionRepository) DeleteRolePermissionByPermissionId(permissionId int) error {
	panic("implement me")
}

func (r *RolePermissionRepository) GetPermissionByRoleId(roleId int) ([]*model.Permission, error) {
	panic("implement me")
}
