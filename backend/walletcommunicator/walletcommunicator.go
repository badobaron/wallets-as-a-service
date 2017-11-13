package walletcommunicator

import (
	"fmt"

	"github.com/blockcypher/gobcy"
)

var bcy = gobcy.API{Token: "2aa27c3912c047f2baa7e932cfc453e7", Coin: "bcy", Chain: "test"}

func GetBalance(address string) int {
	addr, err := bcy.GetAddrBal(address, nil)
	if err != nil {
		fmt.Println(err)
	}

	return addr.Balance
}

func CreateWallet() gobcy.AddrKeychain {
	addrKeys, err := bcy.GenAddrKeychain()
	if err != nil {
		fmt.Println(err)
	}
	return addrKeys
}

func SendTransaction(source string, dest string, amount int) {
	//Post New TXSkeleton
	skel, err := bcy.NewTX(gobcy.TempNewTX(source, dest, 1000), false)
	//Sign it locally
	err = skel.Sign([]string{"0e6372d9fc09b0b345ed4a8f9477d0b12c6c5b1ff7f352c4a53cf79ee3d10f06"})
	if err != nil {
		fmt.Println(err)
	}
	//Send TXSkeleton
	skel, err = bcy.SendTX(skel)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", skel)
}

func FaucetToAddress(address string) {
	pair := gobcy.AddrKeychain{}
	pair.Address = address
	txhash, err := bcy.Faucet(pair, 100000)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(txhash)
}
