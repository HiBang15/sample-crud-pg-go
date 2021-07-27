package ipm

import (
	"github.com/HiBang15/sample-crud-pg-go.git/model"
	"github.com/HiBang15/sample-crud-pg-go.git/repository"
	"github.com/go-pg/pg"
)

type UserRoleRepository struct {
	Connector *pg.DB
}

func NewUserRoleRepository(db *pg.DB) repository.IUserRoleRepository {
	return &UserRoleRepository{Connector: db}
}

func (u *UserRoleRepository) GetListUserRole() ([]*model.UserRole, error) {
	panic("implement me")
}

func (u *UserRoleRepository) GetUserRoleByRoleId(roleId int) (*model.UserRole, error) {
	panic("implement me")
}

func (u *UserRoleRepository) GetUserRoleByUserId(userId int) (*model.UserRole, error) {
	panic("implement me")
}

func (u *UserRoleRepository) CreateUserRole(role *model.UserRole) error {
	panic("implement me")
}

func (u *UserRoleRepository) UpdateUserRoleByUserId(role *model.UserRole) error {
	panic("implement me")
}

func (u *UserRoleRepository) UpdateUserRoleByRoleId(role *model.UserRole) error {
	panic("implement me")
}

func (u *UserRoleRepository) DeleteUserRoleByUserId(userID int) error {
	panic("implement me")
}

func (u *UserRoleRepository) DeleteUserRoleByRoleId(roleId int) error {
	panic("implement me")
}

func (u *UserRoleRepository) GetRoleByUserId(userId int) ([]*model.Role, error) {
	panic("implement me")
}
