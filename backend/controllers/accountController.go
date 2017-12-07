package controllers

import (
	"net/http"
	"github.com/wandi34/wallets-as-a-service/backend/data"
	"github.com/blockcypher/gobcy"
	"fmt"
	"encoding/json"
	"github.com/wandi34/wallets-as-a-service/backend/common"
	"github.com/wandi34/wallets-as-a-service/backend/models"
	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/mux"
	"crypto/aes"
	"bytes"
)


var bcy = gobcy.API{"2aa27c3912c047f2baa7e932cfc453e7", "bcy", "test"}

func GetAccounts(w http.ResponseWriter, r *http.Request) {
	// Get userId from URL
	vars := mux.Vars(r)
	userId := bson.ObjectIdHex(vars["id"])

	// Get Account for userId
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("accounts")
	repo := &data.AccountRepository{C: col}
	// Authenticate the login user
	result := models.Account{}
	err := repo.C.Find(bson.M{"userid": userId}).One(&result)
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
	w.WriteHeader(http.StatusOK)
	j, err := json.Marshal(AccountResource{Data: result})
	w.Write(j)

}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var dataResource CreateWalletResource
	// Decode the incoming CreateAccount json
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
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("accounts")
	repo := &data.AccountRepository{C: col}
	// Create new wallet
	addrKeys := createAddress()
	// Encrypt private key with guard
	addrKeys.Private = encryptPrivateKey(addrKeys.Private, "1234")
	// Insert account document
	repo.CreateAccount(addrKeys, dataResource.Data.UserId)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func createAddress() gobcy.AddrKeychain {
	addrKeys, err := bcy.GenAddrKeychain()
	if err != nil {
		fmt.Println(err)
	}
	return addrKeys
}

func encryptPrivateKey(privateKey, secret string) string {
	bc, err := aes.NewCipher([]byte(secret))
	if (err != nil) {
		fmt.Println(err)
	}

	var dst = make([]byte, 16)
	var src = []byte(privateKey)

	bc.Encrypt(dst, src)
	return bytes.NewBuffer(dst).String()
}


