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
	//TODO: Encrypt key with userCredentialGuard
	// Encrypt private key with guard
	password := dataResource.Data.Password
	encryptedBytes, _ := common.Encrypt(password, addrKeys.Private, dataResource.Data.UserId.String())
	addrKeys.Private = string(encryptedBytes[:])
	// Create iban for cryptowallet
	iban := createIBAN()
	// Insert account document
	repo.CreateAccount(addrKeys, dataResource.Data.UserId, common.GetMd5Hash(password), iban)

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

func createIBAN() string {
	// For real implementation get a free iban address
	// For prototype, iban will be created from the service
	return "DE89 3704 0044 0532 0130 00"
}


