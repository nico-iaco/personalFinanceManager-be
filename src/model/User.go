package model

type User struct {
	ID          string   `json:"id"`
	Enabled     bool     `json:"isEnabled" default:"false"`
	Firstname   string   `json:"firstname"`
	Lastname    string   `json:"lastname"`
	Email       string   `json:"email"`
	Password    string   `json:"-"`
	Accounts    []string `json:"accounts"`
	Investments []string `json:"investments"`
	Cryptos     []string `json:"cryptos"`
}
