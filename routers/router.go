package routers

import (
	"final-project/controller"

	"github.com/gin-gonic/gin"
)

type Router struct {
	router         *gin.Engine
	userController controller.UserController
	port           string
}

func NewRouter(gin *gin.Engine, userController controller.UserController, port string) *Router {
	return &Router{
		router:         gin,
		userController: userController,
		port:           port,
	}
}

func (r *Router) StartServer() {
	r.router.POST("/users", r.userController.RegisterUser)
	r.router.GET("/users", r.userController.GetAllUser)
	r.router.Run(r.port)
}
