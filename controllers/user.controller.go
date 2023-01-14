package controllers

import (
	"crude-app/models"
	"crude-app/services"
	"net/http"

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

	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := uc.UserService.CreateUser(&user)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	username := ctx.Param("name")
	user, err := uc.UserService.GetUser(&username)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)

}

func (uc *UserController) GetAll(ctx *gin.Context) {
	user, err := uc.UserService.GetAll()

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {

	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := uc.UserService.UpdateUser(&user)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	username := ctx.Param("name")
	err := uc.UserService.DeleteUser(&username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "record deleted"})
}

// Routes
func (uc *UserController) RegisterUserRoute(rg *gin.RouterGroup) {
	userroute := rg.Group("/user")

	userroute.POST("/create", uc.CreateUser)
	userroute.GET("/get/:name", uc.GetUser)
	userroute.GET("/getAll", uc.GetAll)
	userroute.POST("/update", uc.UpdateUser)
	userroute.DELETE("/delete/:name", uc.CreateUser)

}
