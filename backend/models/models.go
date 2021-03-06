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
		UserId		bson.ObjectId	`json:"userId"`
		Iban		string			`json:"iban"`
		PwHash		[]byte			`json:"pwHash"`
		Wallet		gobcy.AddrKeychain	`json:"addrKeys"`
	}
	Iban struct {
		TargetAddress		string			`json:"targetAddress"`
		Amount				string			`json:"amount"`
	}
)
