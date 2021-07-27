package ipm

import (
	"github.com/HiBang15/sample-crud-pg-go.git/model"
	"github.com/HiBang15/sample-crud-pg-go.git/repository"
	"github.com/go-pg/pg"
)

type PermissionRepository struct {
	Connector *pg.DB
}

func NewPermissionRepository(db *pg.DB) repository.IPermissionRepository {
	return &PermissionRepository{Connector: db}
}

func (p *PermissionRepository) GetListPermission() ([]*model.Permission, error) {
	panic("implement me")
}

func (p *PermissionRepository) GetPermissionById(permissionId int) (*model.Permission, error) {
	panic("implement me")
}

func (p *PermissionRepository) CreatePermission(permission *model.Permission) error {
	panic("implement me")
}

func (p *PermissionRepository) UpdatePermission(permission *model.Permission) error {
	panic("implement me")
}

func (p *PermissionRepository) DeletePermission(permissionId int) error {
	panic("implement me")
}
