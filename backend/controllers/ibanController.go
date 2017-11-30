package controllers

import (
	"net/http"
	"encoding/json"
	"github.com/wandi34/wallets-as-a-service/backend/common"
	"fmt"
	"strings"
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
	var amount float64
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


	_, err = bcy.Faucet("CCUg986pdG5J3ek52kyMs9XhF5TqV8ND8C", 3e5)
	if err != nil {
		fmt.Println(err)
	}


}
