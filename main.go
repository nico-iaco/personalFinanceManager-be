package main

import (
	"github.com/gin-gonic/gin"
	"personalFinanceManager/controller"
	"personalFinanceManager/middleware"
	"personalFinanceManager/repository"
)

func main() {
	r := gin.Default()
	repository.CreateConnection()
	defer repository.Disconnect()

	r.POST("/register", controller.RegisterUser) //Api to register user
	r.POST("/login", controller.Login)           //Api to login user

	accountController := r.Group("/user/account").Use(middleware.AuthorizeJwt())
	{
		accountController.POST("", controller.AddAccount)
		accountController.PATCH("", controller.EditAccount)
		accountController.DELETE("", controller.DeleteAccount)
	}

	err := r.Run()
	if err != nil {
		return
	}
}
