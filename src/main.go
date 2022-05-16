package main

import (
	"github.com/gin-gonic/gin"
	"personalFinanceManager/src/controller"
	"personalFinanceManager/src/middleware"
	"personalFinanceManager/src/repository"
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

	accountMovementsController := r.Group("/user/account/movement").Use(middleware.AuthorizeJwt())
	{
		accountMovementsController.GET("", controller.GetUserAccountMovements)
		accountMovementsController.POST("", controller.AddAccountMovement)
		accountMovementsController.DELETE("/:movementId", controller.DeleteAccountMovement)
	}

	err := r.Run()
	if err != nil {
		return
	}
}
