package walletcommunicator

import (
	"io/ioutil"
	"net/http"
)

func GetBalance(apikey string) string {
	rs, err := http.Get("https://block.io/api/v2/get_balance/?api_key=" + apikey)

	if err != nil {
		panic(err) // More idiomatic way would be to print the error and die unless it's a serious error
	}
	defer rs.Body.Close()

	bodyBytes, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		panic(err)
	}

	bodyString := string(bodyBytes)

	return bodyString
}
