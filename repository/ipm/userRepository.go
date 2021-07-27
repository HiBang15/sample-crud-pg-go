package ipm

import (
	"github.com/HiBang15/sample-crud-pg-go.git/model"
	"github.com/HiBang15/sample-crud-pg-go.git/repository"
	"github.com/go-pg/pg"
	"log"
)

type UserRepository struct {
	Connector *pg.DB
}

func NewUserRepository(db *pg.DB) repository.IUserRepository {
	return &UserRepository{Connector: db}
}

func (u *UserRepository) GetUserByEmail(email string) (user *model.User, err error) {
	err = u.Connector.Model(user).Where("email = ?", email).Select()
	if err != nil {
		log.Println("[UserRepository][Error]Get user by email fail with err: " + err.Error())
		return nil, err
	}

	log.Println("[UserRepository][Success]Get user by email successful")
	return user, nil
}

func (u *UserRepository) GetListUser() (users []*model.User, err error) {
	err = u.Connector.Model(users).Select()
	if err != nil {
		log.Println("[UserRepository][Error]Get list user fail with err: " + err.Error())
		return nil, err
	}

	log.Println("[UserRepository][Success]Get list user successful!")
	return users, nil
}

func (u *UserRepository) GetUserById(userID int) (user *model.User, err error) {
	err = u.Connector.Model(user).Where("id = ?", userID).Select()
	if err != nil {
		log.Println("[UserRepository][Error]Get user by id fail with err: " + err.Error())
		return nil, err
	}

	log.Println("[UserRepository][Success]Get user by id successful")
	return user, nil
}

func (u *UserRepository) CreateUser(user *model.User) ( err error) {
	_, err = u.Connector.Model(user).Insert()
	if err != nil {
		log.Println("[UserRepository][Error]Create user fail with err: " + err.Error())
		return err
	}

	log.Println("[UserRepository][Success]Create user successful!")
	return nil
}

func (u *UserRepository) UpdateUser(user *model.User) error {
	//Update bảng user
	_, err := u.Connector.Model(user).TableExpr("auth.users").Where("id = ?", user.Id).Update()

	if err != nil {
		log.Println("[UserRepository][Error]Update user fail with err: " + err.Error())
		return err
	}

	log.Println("[UserRepository][Success] update user successful")
	return nil
}

func (u *UserRepository) DeleteUser(userID int) error {
	user := new(model.User)

	//Xóa user với id
	_, err := u.Connector.Model(user).Where("id = ?", userID).Delete()
	if err != nil {
		log.Println("[UserRepository][Error] Delete user by id fail with err: " + err.Error())
		return err
	}

	log.Println("[UserRepository][Success] delete user successful")
	return nil
}
