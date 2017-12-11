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

func (r *AccountRepository) CreateAccount(addrKeys gobcy.AddrKeychain, userId bson.ObjectId) error {
	account := models.Account{bson.NewObjectId(), userId, addrKeys}
	err := r.C.Insert(&account)
	return err
}

func (r *AccountRepository) GetKeyChainFromAddress(address string) (gobcy.AddrKeychain, error) {
	result := models.Account{}
	err := r.C.Find(bson.M{"wallet.address": address}).One(&result)

	return result.Wallet, err
}
