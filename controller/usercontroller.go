package controller

import (
	"final-project/params"
	"final-project/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	db          *gorm.DB
	userService *services.UserServices
}

func NewUserController(db *gorm.DB, userService *services.UserServices) *UserController {
	return &UserController{
		db:          db,
		userService: userService,
	}
}

func (c *UserController) RegisterUser(ctx *gin.Context) {
	var req params.CreateUserRequest

	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	createUser := c.userService.RegisterUser(&req)

	ctx.JSON(createUser.Status, createUser)
}

func (c *UserController) GetAllUser(ctx *gin.Context) {
	getUser := c.userService.GetAllUser()
	ctx.JSON(getUser.Status, getUser)
}
