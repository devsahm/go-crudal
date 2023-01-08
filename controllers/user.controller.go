package controllers

import (
	"crude-app/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func New(UserService services.UserService) UserController {
	return UserController{
		UserService: UserService,
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	ctx.JSON(200, " ")
}

// func (u *UserServiceImpl) GetUser(name *string) {
// 	return nil, nil
// }

// func (u *UserServiceImpl) GetAll() {
// 	return nil, nil
// }

// func (u *UserServiceImpl) UpdateUser(user *models.User)  {
// 	return nil
// }

// func (u *UserServiceImpl) DeleteUser(name *string)  {
// 	return nil
// }
