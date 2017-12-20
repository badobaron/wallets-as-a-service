package controllers

import (
	"net/http"
	"encoding/json"
	"github.com/wandi34/wallets-as-a-service/backend/common"
	"github.com/wandi34/wallets-as-a-service/backend/models"
	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/mux"
	"github.com/wandi34/wallets-as-a-service/backend/data"
)

func GetAccounts(w http.ResponseWriter, r *http.Request) {
	// Get userId from URL
	vars := mux.Vars(r)
	userId := bson.ObjectIdHex(vars["id"])

	// Get Account for userId
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("accounts")
	repo := &data.AccountRepository{C: col}

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
		common.DisplayAppError(w, err, "Invalid body", 500, )
		return
	}
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("accounts")
	repo := &data.AccountRepository{C: col}
	// Create new wallet
	addrKeys, err := common.CreateAddress()
	if err != nil {
		common.DisplayAppError(w, err, "Account creation error", 500, )
		return
	}
	// Split private key into parts with SSSS for restore function
	common.SplitSecret(addrKeys.Private)
	// Encrypt private key with password
	password := dataResource.Data.Password
	encryptedBytes, _ := common.Encrypt(password, addrKeys.Private, dataResource.Data.UserId.String())
	addrKeys.Private = string(encryptedBytes[:])
	// Create iban for cryptowallet
	iban := common.CreateIBAN()
	// Insert account document
	repo.CreateAccount(addrKeys, dataResource.Data.UserId, common.GetMd5Hash(password), iban)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

//func getAccountRepository() *data.AccountRepository {
//	context := NewContext()
//	defer context.Close()
//	col := context.DbCollection("accounts")
//	repo := &data.AccountRepository{C: col}
//	return repo
//}

