package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"personalFinanceManager/utils"
)

func AuthorizeJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := utils.DecodeJwt(tokenString)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			userId := claims["userId"]
			c.Set("userId", userId)
			fmt.Println(claims)
		} else {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
