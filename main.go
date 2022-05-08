package main

import (
	"github.com/gin-gonic/gin"
	"personalFinanceManager/controller"
)

func main() {
	r := gin.Default()

	r.POST("/user/register", controller.RegisterUser) //Api to register user
	r.POST("/user/login", controller.Login)           //Api to login user

	err := r.Run()
	if err != nil {
		return
	}
}
