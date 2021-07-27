package ipm

import (
	"github.com/HiBang15/sample-crud-pg-go.git/model"
	"github.com/HiBang15/sample-crud-pg-go.git/repository"
	"github.com/go-pg/pg"
)

type PermissionActivityRepository struct {
	Connector *pg.DB
}

func NewPermissionActivityRepository(db *pg.DB) repository.IPermissionActivityRepository {
	return &PermissionActivityRepository{Connector: db}
}

func (p *PermissionActivityRepository) GetListPermissionActivity() ([]*model.PermissionActivity, error) {
	panic("implement me")
}

func (p *PermissionActivityRepository) GetPermissionActivityByPermissionId(permissionId int) (*model.PermissionActivity, error) {
	panic("implement me")
}

func (p *PermissionActivityRepository) CreatePermissionActivity(activity *model.PermissionActivity) error {
	panic("implement me")
}

func (p *PermissionActivityRepository) UpdatePermissionActivityByPermissionId(activity *model.PermissionActivity) error {
	panic("implement me")
}

func (p *PermissionActivityRepository) DeletePermissionActivityByPermissionId(permissionId int) error {
	panic("implement me")
}

func (p *PermissionActivityRepository) GetActivityByPermissionId(permissionId int) ([]*model.Permission, error) {
	panic("implement me")
}
