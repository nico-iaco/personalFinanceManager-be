package response

import "personalFinanceManager/model"

type LoginResponse struct {
	User model.User `json:"user"`
	Jwt  string     `json:"jwt"`
}
