package controllers

import (
	"net/http"
	_ "github.com/wandi34/wallets-as-a-service/backend/data"
	"github.com/wandi34/wallets-as-a-service/backend/data"
	"github.com/blockcypher/gobcy"
	"fmt"
	"encoding/json"
	"github.com/wandi34/wallets-as-a-service/backend/common"
)


var bcy = gobcy.API{"2aa27c3912c047f2baa7e932cfc453e7", "bcy", "test"}

func GetAccounts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var dataResource CreateWalletResource
	// Decode the incoming User json
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


