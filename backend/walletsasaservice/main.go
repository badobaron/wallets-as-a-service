package main

import (
	"wallets-as-a-service/backend/walletcommunicator"
)

func main() {
	source := "C5sjnsYL4DS1vMNEGyFCzYaFDeJfRneHur"
	dest := "CFTMcVri441k51yhzc3UbNmE2c2iNtVorw"
	// fmt.Println(walletcommunicator.GetBalance("C5sjnsYL4DS1vMNEGyFCzYaFDeJfRneHur"))
	// fmt.Println(walletcommunicator.GetBalance("CFTMcVri441k51yhzc3UbNmE2c2iNtVorw"))
	// // walletcommunicator.FaucetToAddress("C5sjnsYL4DS1vMNEGyFCzYaFDeJfRneHur")
	// fmt.Println(walletcommunicator.GetBalance("C5sjnsYL4DS1vMNEGyFCzYaFDeJfRneHur"))
	// fmt.Println(w)

	walletcommunicator.SendTransaction(source, dest, 1000)
}
