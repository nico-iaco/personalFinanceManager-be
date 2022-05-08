package model

// Movement
// this class represent an economic movement if source is populated the movement is an expense
// otherwise is an income, if either source and destination is populated the movement is a transfer between two account
// or is a recurring investment from an account to an investment
///**
type Movement struct {
	ID          uint    `json:"id"`
	Source      string  `json:"source,omitempty"`
	Destination string  `json:"destination,omitempty"`
	Amount      float64 `json:"amount"`
	Note        string  `json:"note,omitempty"`
}
