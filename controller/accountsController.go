package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"personalFinanceManager/model"
	"personalFinanceManager/model/response"
	"personalFinanceManager/repository/user"
)

func AddAccount(c *gin.Context) {
	userId := c.GetString("userId")
	accountName := c.Query("accountName")
	log.Printf("Adding %v account to user %v", accountName, userId)
	u := user.GetUserById(userId)
	accounts := u.Accounts
	if accounts == nil {
		accountsArray := []string{accountName}
		u.Accounts = accountsArray
	} else {
		u.Accounts = append(accounts, accountName)
	}
	updatedFields := bson.D{
		{"$set", bson.D{
			{"accounts", u.Accounts},
		}},
	}
	user.UpdateUser(u, updatedFields)
	c.JSON(http.StatusOK, response.BaseResponse[model.User]{
		http.StatusOK,
		u,
		"",
	})
	return
}

func EditAccount(c *gin.Context) {
	userId := c.GetString("userId")
	oldName := c.Query("oldName")
	newName := c.Query("newName")
	u := user.GetUserById(userId)
	accounts := u.Accounts
	if accounts == nil {
		c.JSON(http.StatusOK, response.BaseResponse[string]{
			http.StatusBadRequest,
			"",
			"Logged user hasn't any accounts added",
		})
		return
	} else {
		indexOfString := funk.IndexOfString(accounts, oldName)
		accounts[indexOfString] = newName
		u.Accounts = accounts
	}
	updatedFields := bson.D{
		{"$set", bson.D{
			{"accounts", u.Accounts},
		}},
	}
	user.UpdateUser(u, updatedFields)
	c.JSON(http.StatusOK, response.BaseResponse[model.User]{
		http.StatusOK,
		u,
		"",
	})
	return
}

func DeleteAccount(c *gin.Context) {
	userId := c.GetString("userId")
	accountName := c.Query("accountName")
	u := user.GetUserById(userId)
	accounts := u.Accounts
	if accounts == nil {
		c.JSON(http.StatusOK, response.BaseResponse[string]{
			http.StatusBadRequest,
			"",
			"Logged user hasn't any accounts added",
		})
		return
	} else {
		u.Accounts = funk.FilterString(accounts, func(s string) bool {
			return s != accountName
		})
	}
	updatedFields := bson.D{
		{"$set", bson.D{
			{"accounts", u.Accounts},
		}},
	}
	user.UpdateUser(u, updatedFields)
	c.JSON(http.StatusOK, response.BaseResponse[model.User]{
		http.StatusOK,
		u,
		"",
	})
	return
}
