package models

import (

	"gopkg.in/mgo.v2/bson"
	"github.com/blockcypher/gobcy"
)

type (
	User struct {
		Id           bson.ObjectId `bson:"_id,omitempty" json:"id"`
		FirstName    string        `json:"firstname"`
		LastName     string        `json:"lastname"`
		Email        string        `json:"email"`
		Password     string        `json:"password,omitempty"`
		HashPassword []byte        `json:"hashpassword,omitempty"`
	}
	Account struct {
		Id			bson.ObjectId  `bson:"_id,omitempty" json:"id"`
		Wallet		gobcy.Wallet	`json:"wallet"`
	}
)
