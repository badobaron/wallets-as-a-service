package data

import (
	"gopkg.in/mgo.v2"
	"github.com/wandi34/wallets-as-a-service/backend/models"
	"gopkg.in/mgo.v2/bson"
)

type AccountRepository struct {
	C *mgo.Collection
}

func (r *AccountRepository) CreateAccount(account *models.Account) error {
	obj_id := bson.NewObjectId()
	account.Id = obj_id
	err := r.C.Insert(&account)
	return err
}
