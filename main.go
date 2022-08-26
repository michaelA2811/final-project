package main

import (
	"final-project/config"
	"final-project/controller"
	"final-project/repository"
	"final-project/routers"
	"final-project/services"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func main() {
	config.LoadConfig("config")
	db, err := config.Connect()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	port := ":" + viper.GetString("server.port")

	userRepo := repository.NewUserRepository(db)
	userService := services.NewCarServices(userRepo)
	userController := controller.NewUserController(db, userService)

	router := routers.NewRouter(r, *userController, port)

	router.StartServer()

}
