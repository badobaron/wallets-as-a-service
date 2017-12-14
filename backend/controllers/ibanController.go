package controllers

import (
	"net/http"
	"encoding/json"
	"github.com/wandi34/wallets-as-a-service/backend/common"
	"strings"
	"github.com/wandi34/wallets-as-a-service/backend/data"
)

func ConvertIban(w http.ResponseWriter, r *http.Request) {
	var dataResource IbanResource
	// Decode the incoming Iban json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid User data",
			500,
		)
		return
	}
	ibanTransaction := &dataResource.Data
	//Convert euro in cryptocurrency
	url := "https://blockchain.info/tobtc?currency=EUR&value=" + strings.Replace(ibanTransaction.Amount, ",",".", -1)
	resp, err := http.Get(url)
	var amount float32
	err = json.NewDecoder(resp.Body).Decode(&amount)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid User data",
			500,
		)
		return
	}

	context := NewContext()
	defer context.Close()
	col := context.DbCollection("accounts")
	repo := &data.AccountRepository{C: col}

	address, err := repo.GetAddressFromIban(dataResource.Data.TargetAddress)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"IBAN not found",
			400,
		)
		return
	}
	txHash, err := faucet(&w, address, int(amount))
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Faucet error",
			500,
		)
		return
	}
	j, err := json.Marshal(txHash)
	w.Write(j)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}

func faucet(w *http.ResponseWriter, address string, amount int) (string, error) {
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("accounts")
	repo := &data.AccountRepository{C: col}
	keyChain, err := repo.GetKeyChainFromAddress(address)
	if err != nil {
			common.DisplayAppError(
				*w,
				err,
				"No such address",
				500,
			)
	}

	return bcy.Faucet(keyChain, amount)
}
