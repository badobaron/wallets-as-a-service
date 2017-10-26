package main

import (
	"log"
	"net/http"
	"io/ioutil"
)

func main() {
	apikey := "45cf-7687-926e-1ae2"
	
//	log.Println(bodyString)
	
	balance string = walletcommunicator.getBalance(apikey)
	
	log.Println(balance)

}