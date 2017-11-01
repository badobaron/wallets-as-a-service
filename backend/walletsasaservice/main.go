package main

import (
	"log"

	"wallets-as-a-service/backend/walletcommunicator"
)

func main() {
	apikey := "45cf-7687-926e-1ae2"

	var balance string = walletcommunicator.GetBalance(apikey)

	log.Println(balance)
}
