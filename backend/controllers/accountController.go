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


	addrKeys, err := bcy.GenAddrMultisig(gobcy.AddrKeychain{PubKeys: []string{"02c716d071a76cbf0d29c29cacfec76e0ef8116b37389fb7a3e76d6d32cf59f4d3", "033ef4d5165637d99b673bcdbb7ead359cee6afd7aaf78d3da9d2392ee4102c8ea", "022b8934cc41e76cb4286b9f3ed57e2d27798395b04dd23711981a77dc216df8ca"}, ScriptType: "multisig-2-of-3"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", addrKeys)
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("accounts")
	repo := &data.AccountRepository{C: col}
	// Insert account document

	repo.CreateAccount()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}


