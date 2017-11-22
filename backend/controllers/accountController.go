package controllers

import (
	"net/http"
	_ "github.com/wandi34/wallets-as-a-service/backend/data"
	"github.com/wandi34/wallets-as-a-service/backend/data"
	"github.com/blockcypher/gobcy"
	"fmt"
)


var bcy = gobcy.API{"2aa27c3912c047f2baa7e932cfc453e7", "bcy", "test"}

func GetAccounts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("accounts")
	repo := &data.AccountRepository{C: col}
	// Create new wallet
	wallet := createWallet("test123321")
	// Insert account document

	repo.CreateAccount(wallet)

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

func createWallet(userId string) gobcy.Wallet{
	keychain := createAddress()
	wallet, err := bcy.CreateWallet(gobcy.Wallet{userId, []string{keychain.Address}})

	if err != nil {
		fmt.Println(err)
	}
	return wallet
}


