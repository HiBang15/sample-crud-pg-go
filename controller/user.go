package controller

import (
	"github.com/HiBang15/sample-crud-pg-go.git/lib"
	"github.com/HiBang15/sample-crud-pg-go.git/model"
	"github.com/HiBang15/sample-crud-pg-go.git/service"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	"log"
	"net/http"
	"strconv"
)

var DB *pg.DB

func CreateUser(ctx *gin.Context)  {
	var user *model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		lib.SetResponse(ctx, http.StatusUnprocessableEntity, err, "Invalid Request Body", nil)
		return
	}

	//create user via user service
	userService := service.NewUserService(DB)
	err := userService.CreateUser(user)
	if err != nil {
		lib.SetResponse(ctx, http.StatusInternalServerError, err, err.Error(), 0)
		return
	}

	lib.SetResponse(ctx, http.StatusOK, nil, "create user successful!", user)
	return
}

func UpdateUserAccount(c *gin.Context) {
	var userInfoUpdate *model.User
	//get id
	id := c.Params.ByName("id")
	user_id, err := strconv.Atoi(id)

	if err != nil {
		log.Printf("Update user account fails with error: %v", err)
		lib.SetResponse(c, http.StatusInternalServerError, err, "Cannot update user", nil)
		return
	}

	if err := c.ShouldBindJSON(&userInfoUpdate); err != nil {
		lib.SetResponse(c, http.StatusUnprocessableEntity, err, "Invalid request body!", nil)
		return
	}
	userInfoUpdate.Id = user_id

	//update user
	userService := service.NewUserService(DB)
	err = userService.UpdateUser(userInfoUpdate)

	if err != nil {
		log.Printf("update user fails with error: %v", err)
		lib.SetResponse(c, http.StatusInternalServerError, err, "Cannot update user", nil)
		return
	}

	lib.SetResponse(c, http.StatusOK, err, "update user successful!", userInfoUpdate)
	return
}
func GetUserById(c *gin.Context) {
	id := c.Params.ByName("id")
	userId, err := strconv.Atoi(id)

	if err != nil {
		log.Printf("get user account with error: %v", err)
		lib.SetResponse(c, http.StatusInternalServerError, err, "Cannot get user", nil)
		return
	}

	//get user account
	userService := service.NewUserService(DB)
	res, err := userService.GetUserByUserId(userId)
	if err != nil {
		log.Printf("get user fails with error: %v", err)
		lib.SetResponse(c, http.StatusInternalServerError, err, "Cannot get user", nil)
		return
	}
	lib.SetResponse(c, http.StatusOK, err, "get user successful!", res)
	return
}

// DeleteUserAccount handler delete user account
// @Security	bearerAuth
// @Description	Delete user account
// @Tags	account
// @Param	id	path	integer	true	"Id of User Account"
// @Param	shop_hash	query	string	false	"Shop Hash Of Store"
// @Success	200
// @Failure	403	{object}	ErrorResponse
// @Router	/user/{id}	[delete]
func DeleteUserAccount(c *gin.Context) {
	id := c.Params.ByName("id")
	//shopHash := c.Query("shop_hash")
	userId, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("delete user fails with error: %v", err)
		lib.SetResponse(c, http.StatusInternalServerError, err, "Cannot delete user", nil)
		return
	}

	userService := service.NewUserService(DB)
	err = userService.DeleteUser(userId)
		if err != nil {
		log.Printf("delete user fails with error: %v", err)
		lib.SetResponse(c, http.StatusInternalServerError, err, err.Error(), false)
		return
	}

	lib.SetResponse(c, http.StatusOK, err, "delete user successful!", true)
	return
}

func GetAllUser(c *gin.Context) {

	userService := service.NewUserService(DB)
	res, err := userService.GetListUser()
	if err != nil {
		lib.SetResponse(c, http.StatusInternalServerError, err, err.Error(), res)
		return
	}
	lib.SetResponse(c, http.StatusOK, err, "Get list user successful", res)
	return
}