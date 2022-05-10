package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Movement
// this class represent an economic movement if source is populated the movement is an expense
// otherwise is an income, if either source and destination is populated the movement is a transfer between two account
// or is a recurring investment from an account to an investment
///**
type Movement struct {
	ID          string             `json:"id"`
	User        string             `json:"userId"`
	Source      string             `json:"source,omitempty"`
	Destination string             `json:"destination,omitempty"`
	Amount      float64            `json:"amount"`
	Date        primitive.DateTime `json:"date"`
	Note        string             `json:"note,omitempty"`
}
