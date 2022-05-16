package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"net/http"
	"personalFinanceManager/src/model"
	"personalFinanceManager/src/model/request"
	"personalFinanceManager/src/model/response"
	"personalFinanceManager/src/repository/user"
	"personalFinanceManager/src/utils"
)

func RegisterUser(c *gin.Context) {
	var input request.RegistrationRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusOK, response.BaseResponse[string]{
			400,
			"",
			err.Error(),
		})
		return
	}

	emailExists := user.CheckEmailExists(input.Email)
	if emailExists {
		c.JSON(http.StatusOK, response.BaseResponse[string]{
			402,
			"",
			"Email already exists",
		})
		return
	}

	encodedPassword, _ := utils.HashPassword(input.Password)

	u := model.User{
		ID:          xid.New().String(),
		Firstname:   input.Firstname,
		Lastname:    input.Lastname,
		Email:       input.Email,
		Password:    encodedPassword,
		Enabled:     false,
		Accounts:    nil,
		Investments: nil,
		Cryptos:     nil,
	}

	userAdded := user.AddUser(u)
	c.JSON(http.StatusOK, response.BaseResponse[model.User]{
		http.StatusOK,
		userAdded,
		"",
	})
	return
}

func Login(c *gin.Context) {
	var input request.LoginRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusOK, response.BaseResponse[string]{
			400,
			"",
			err.Error(),
		})
		return
	}
	u := user.GetUser(input.Email)
	if !utils.CheckPasswordHash(input.Password, u.Password) {
		c.JSON(http.StatusOK, response.BaseResponse[string]{
			400,
			"",
			"Invalid username o password",
		})
		return
	}
	jwt := utils.GenerateJwt(u)
	loginResponse := response.LoginResponse{
		User: u,
		Jwt:  jwt,
	}
	c.JSON(http.StatusOK, response.BaseResponse[response.LoginResponse]{
		http.StatusOK,
		loginResponse,
		"",
	})
	return
}
