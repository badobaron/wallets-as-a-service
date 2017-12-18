package data

import (
	"gopkg.in/mgo.v2"
	"github.com/wandi34/wallets-as-a-service/backend/models"
	"gopkg.in/mgo.v2/bson"
	"github.com/blockcypher/gobcy"
)

type AccountRepository struct {
	C *mgo.Collection
}

func (r *AccountRepository) CreateAccount(addrKeys gobcy.AddrKeychain, userId bson.ObjectId, pwHash []byte, iban string) error {
	account := models.Account{bson.NewObjectId(), userId, iban,pwHash, addrKeys}
	err := r.C.Insert(&account)
	return err
}

func (r *AccountRepository) GetKeyChainFromAddress(address string) (gobcy.AddrKeychain, error) {
	result := models.Account{}
	err := r.C.Find(bson.M{"wallet.address": address}).One(&result)

	return result.Wallet, err
}

func (r *AccountRepository) GetAddressFromIban(iban string) (string, error) {
	result := models.Account{}
	err := r.C.Find(bson.M{"iban": iban}).One(&result)

	return result.Wallet.Address, err
}
