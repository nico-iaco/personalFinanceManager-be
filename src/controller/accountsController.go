package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"github.com/thoas/go-funk"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"personalFinanceManager/src/model"
	"personalFinanceManager/src/model/response"
	"personalFinanceManager/src/repository/movement"
	"personalFinanceManager/src/repository/user"
	"personalFinanceManager/src/utils"
)

func AddAccount(c *gin.Context) {
	userId := c.GetString("userId")
	accountName := c.Query("accountName")
	log.Printf("Adding %v account to user %v", utils.SanitizeString(accountName), utils.SanitizeString(userId))
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

func AddAccountMovement(c *gin.Context) {
	userId := c.GetString("userId")
	var m model.Movement
	var accountName string
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusOK, response.BaseResponse[string]{
			400,
			"",
			err.Error(),
		})
		return
	}
	if m.Source != "" {
		accountName = m.Source
	} else if m.Destination != "" {
		accountName = m.Destination
	}
	u := user.GetUserById(userId)
	accounts := u.Accounts
	_, founded := funk.FindString(accounts, func(s string) bool {
		return s == accountName
	})
	if !founded {
		c.JSON(http.StatusOK, response.BaseResponse[string]{
			http.StatusBadRequest,
			"",
			"Logged user hasn't any accounts with " + accountName + " name",
		})
		return
	}
	m.ID = xid.New().String()
	m.User = u.ID
	addedMovement := movement.AddMovement(m)
	c.JSON(http.StatusOK, response.BaseResponse[model.Movement]{
		http.StatusOK,
		addedMovement,
		"",
	})
}

func DeleteAccountMovement(c *gin.Context) {
	userId := c.GetString("userId")
	movementId := c.Param("movementId")
	log.Printf("Deleting %v movement", utils.SanitizeString(movementId))
	isMovementDeleted := movement.DeleteUserMovement(userId, movementId)
	if !isMovementDeleted {
		c.JSON(http.StatusOK, response.BaseResponse[string]{
			http.StatusBadRequest,
			"",
			"The movement was not deleted",
		})
		return
	}
	c.JSON(http.StatusOK, response.BaseResponse[string]{
		http.StatusOK,
		"",
		"",
	})
}

func GetUserAccountMovements(c *gin.Context) {
	userId := c.GetString("userId")
	accountName := c.Query("accountName")
	movements := movement.GetUserAccountMovements(userId, accountName)
	c.JSON(http.StatusOK, response.BaseResponse[[]*model.Movement]{
		http.StatusOK,
		movements,
		"",
	})
}
