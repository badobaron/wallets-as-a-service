package walletcommunicator

import (
	"io/ioutil"
	"net/http"
)

func GetBalance(address string) string {
	balance, err := http.Get("https://api.blockcypher.com/v1/bcy/test/addrs/" + address + "/balance")

	if err != nil {
		panic(err)
	}

	defer balance.Body.Close()

	bodyBytes, err := ioutil.ReadAll(balance.Body)
	bodyString := string(bodyBytes)
	return bodyString
}
