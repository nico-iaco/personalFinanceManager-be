package model

import "github.com/kamva/mgm/v3"

type User struct {
	mgm.DefaultModel `bson:",inline" json:"-"`
	ID               string   `json:"id"`
	Enabled          bool     `json:"isEnabled" default:"false"`
	Firstname        string   `json:"firstname"`
	Lastname         string   `json:"lastname"`
	Email            string   `json:"email"`
	Password         string   `json:"-"`
	Accounts         []string `json:"accounts"`
	Investments      []string `json:"investments"`
	Cryptos          []string `json:"cryptos"`
}
