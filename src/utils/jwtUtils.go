package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"personalFinanceManager/src/model"
	"time"
)

var hmacSecret = []byte("test") //FIXME

type authCustomClaims struct {
	Id string `json:"userId"`
	jwt.StandardClaims
}

func GenerateJwt(user model.User) string {
	claims := &authCustomClaims{
		user.ID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    "personal-finance-manager",
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString(hmacSecret)
	if err != nil {
		panic(err)
	}
	return t
}

func DecodeJwt(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSecret, nil
	})

	return token, err
}
