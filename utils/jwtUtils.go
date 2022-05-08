package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
	"personalFinanceManager/model"
)

var hmacSecret = []byte("test") //FIXME

func GenerateJwt(user model.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSecret)
	if err != nil {
		log.Fatal(err.Error())
	}
	return tokenString
}

func DecodeJwt(tokenString string) string {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return fmt.Sprintf("%v", claims["userId"])
	} else {
		fmt.Println(err)
	}
	return ""
}
