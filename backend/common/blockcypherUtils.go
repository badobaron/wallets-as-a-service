package common

import "github.com/blockcypher/gobcy"

var bcy = gobcy.API{"2aa27c3912c047f2baa7e932cfc453e7", "bcy", "test"}

func CreateAddress() (gobcy.AddrKeychain, error) {
	addrKeys, err := bcy.GenAddrKeychain()
	if err != nil {
		return addrKeys, err
	}
	return addrKeys, nil
}

func GetAddress(addressId string) (gobcy.Addr, error){
	addr, err := bcy.GetAddr(addressId, nil)
	return addr, err
}

func FaucetAddress(keychain gobcy.AddrKeychain, amount int) (string, error){
	return bcy.Faucet(keychain, amount)
}

func CreateTransaction(sourceAddress, targetAddress string, amount int) (gobcy.TXSkel, error){
	return bcy.NewTX(gobcy.TempNewTX(sourceAddress, targetAddress, amount), false)
}

func SendTransaction(skel gobcy.TXSkel) (gobcy.TXSkel, error){
	return bcy.SendTX(skel)
}

