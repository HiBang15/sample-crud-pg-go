package ipm

import (
	"github.com/HiBang15/sample-crud-pg-go.git/lib"
	"github.com/HiBang15/sample-crud-pg-go.git/model"
	"github.com/HiBang15/sample-crud-pg-go.git/repository"
	"github.com/go-pg/pg"
)

type RoleRepository struct {
	Connector *pg.DB
}

const fileName = "RoleRepository"

func NewRoleRepository(db *pg.DB) repository.IRoleRepository {
	return &RoleRepository{Connector: db}
}

func (r *RoleRepository) GetListRole() (listRole []*model.Role, err error) {
	err = r.Connector.Model(listRole).Select()
	if err != nil {
		lib.WriteLogError(fileName, "GetListRole", err)
		return nil, err
	}

	lib.WriteLogSuccess(fileName, "GetListRole")
	return listRole, err
}

func (r *RoleRepository) GetRoleById(roleId int) (*model.Role, error) {
	panic("implement me")
}

func (r *RoleRepository) CreateRole(role *model.Role) error {
	panic("implement me")
}

func (r *RoleRepository) UpdateRole(role *model.Role) error {
	panic("implement me")
}

func (r *RoleRepository) DeleteRole(roleId int) error {
	panic("implement me")
}
