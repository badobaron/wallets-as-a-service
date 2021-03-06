package controllers

import (
	"github.com/wandi34/wallets-as-a-service/backend/models"
	"gopkg.in/mgo.v2/bson"
	"github.com/blockcypher/gobcy"
)

//Models for JSON resources
type (
	//For Post - /user/register
	UserResource struct {
		Data models.User `json:"data"`
	}
	//For Post - /iban
	IbanResource struct {
		Data models.Iban `json:"data"`
	}
	//For Post - /wallet/create
	AccountResource struct {
		Data models.Account `json:"data"`
	}
	//For Post - /users/{id}/accounts
	CreateWalletResource struct {
		Data CreateWalletModel `json:"data"`
	}
	//For Post - /users/{id}/accounts/{id}/transactions
	CreateTransactionResource struct {
		Data CreateTransactionModel `json:"data"`
	}
	//For Post - /restore
	RestorePrivateKeyResource struct {
		Data RestorePrivateKeyModel `json:"data"`
	}
	//For Get - /addresses/{id}
	GetAddressResource struct {
		Data gobcy.Addr `json:"data"`
	}
	//For Post - /user/login
	LoginResource struct {
		Data LoginModel `json:"data"`
	}
	//Response for authorized user Post - /user/login
	AuthUserResource struct {
		Data AuthUserModel `json:"data"`
	}
	//Model for authentication
	LoginModel struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	//Model for authorized user with access token
	AuthUserModel struct {
		User  models.User `json:"user"`
		Token string      `json:"token"`
	}
	//Model for creating a wallet
	CreateWalletModel struct {
		UserId		bson.ObjectId 	`json:"userId"`
		Password	string			`json:"password"`
	}

	//Model for creating a wallet
	CreateTransactionModel struct {
		SourceAddress	string `json:"source"`
		TargetAddress	string `json:"target"`
		Amount			string `json:"amount"`
		Password		string `json:"password"`
	}
	//Model for restoring a private key with shares
	RestorePrivateKeyModel struct {
		AccountId		bson.ObjectId	`json:"accountId"`
		Parts			[]string		`json:"parts"`
		Password		string			`json:"password"`
	}


)
