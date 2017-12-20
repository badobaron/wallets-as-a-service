package controllers

import (
	"net/http"
	"encoding/json"
	"github.com/wandi34/wallets-as-a-service/backend/common"
	"github.com/wandi34/wallets-as-a-service/backend/models"
	"gopkg.in/mgo.v2/bson"
	"github.com/SSSaaS/sssa-golang"
)

func RestorePrivateKey(w http.ResponseWriter, r *http.Request) {
	var dataResource RestorePrivateKeyResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid body",
			500,
		)
		return
	}
	accountId := dataResource.Data.AccountId
	password := dataResource.Data.Password
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("accounts")
	// Get requestet account
	result := models.Account{}
	err = col.FindId(accountId).One(&result)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	addrKeys := result.Wallet
	// Decrypt private key with secret parts
	privateKey, err := sssa.Combine(dataResource.Data.Parts)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	// Encrypt private key with guard
	encryptedBytes, _ := common.Encrypt(password, privateKey, result.UserId.String())
	addrKeys.Private = string(encryptedBytes[:])
	// Update document

	err = col.Update(bson.M{"_id": accountId},
		bson.M{"$set": bson.M{
			"wallet": addrKeys,
		}})

	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
