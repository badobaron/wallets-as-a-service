package data

import (
	"gopkg.in/mgo.v2"
	"github.com/wandi34/wallets-as-a-service/backend/models"
	"gopkg.in/mgo.v2/bson"
)

type WalletRepository struct {
	C *mgo.Collection
}

func (r *UserRepository) CreateWallet(wallet *models.Wallet) error {
	obj_id := bson.NewObjectId()
	wallet.Id = obj_id
	err := r.C.Insert(&wallet)
	return err
}
