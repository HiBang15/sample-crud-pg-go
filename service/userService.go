package service

import (
	"errors"
	"fmt"
	"github.com/HiBang15/sample-crud-pg-go.git/model"
	"github.com/HiBang15/sample-crud-pg-go.git/repository"
	"github.com/HiBang15/sample-crud-pg-go.git/repository/ipm"
	"github.com/go-pg/pg"
	"log"
	"regexp"
)

type UserService struct {
	UserRepo repository.IUserRepository
}

func NewUserService(db *pg.DB) *UserService {
	return &UserService{UserRepo: ipm.NewUserRepository(db)}
}

func (userService *UserService) GetListUser() (users []*model.User, err error) {
	log.Println("receive get list user")
	users, err = userService.UserRepo.GetListUser()
	if err != nil {
		return nil, err
	}

	return users, err
}

func (userService *UserService) GetUserByUserId(id int) (user *model.User, err error) {
	log.Println("receive get user with id: %d" + fmt.Sprintf("%d", id))
	if id < 1 {
		log.Println("[UserService][Error]Invalid user id!")
		return nil, errors.New("Invalid user id!")
	}

	user, err = userService.UserRepo.GetUserById(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (userService *UserService)CreateUser(user *model.User) (err error)  {
	log.Println("receive create user with info: ",  user)

	if !isEmailValid(user.Email) {
		log.Println("Invalid Email!")
		return errors.New("Invalid Email!")
	}

	err = userService.UserRepo.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (userService *UserService)UpdateUser(user *model.User) (err error)  {
	log.Println("receive update user with info: ",  user)

	if !isEmailValid(user.Email) {
		log.Println("Invalid Email!")
		return errors.New("Invalid Email!")
	}

	err = userService.UserRepo.UpdateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (userService *UserService) DeleteUser(id int) (err error) {
	log.Println("receive delete user with id: %d" + fmt.Sprintf("%d", id))
	if id < 1 {
		log.Println("[UserService][Error]Invalid user id!")
		return errors.New("Invalid user id!")
	}

	err = userService.UserRepo.DeleteUser(id)
	if err != nil {
		return  err
	}

	return  nil
}

// isEmailValid checks if the email provided is valid by regex.
func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(e)
}
