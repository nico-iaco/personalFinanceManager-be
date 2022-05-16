package response

import (
	"personalFinanceManager/src/model"
)

type LoginResponse struct {
	User model.User `json:"user"`
	Jwt  string     `json:"jwt"`
}
