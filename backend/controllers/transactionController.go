package controllers

import (
	"net/http"
	"encoding/json"
	"github.com/wandi34/wallets-as-a-service/backend/common"
	"fmt"
	"github.com/blockcypher/gobcy"
	"github.com/wandi34/wallets-as-a-service/backend/data"
	"github.com/wandi34/wallets-as-a-service/backend/models"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

//var bcy = gobcy.API{"2aa27c3912c047f2baa7e932cfc453e7", "bcy", "test"}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var dataResource CreateTransactionResource
	// Decode the incoming Transaction json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid body",
			500,
		)
		return
	}
	sourceAddress := dataResource.Data.SourceAddress
	targetAddress := dataResource.Data.TargetAddress
	amount, err := strconv.Atoi(dataResource.Data.Amount)
	//Post New TXSkeleton
	skel, err := bcy.NewTX(gobcy.TempNewTX(sourceAddress, targetAddress, amount), false)
	if err != nil {
		fmt.Println(err)
	}
	//Sign it locally
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("accounts")
	repo := &data.AccountRepository{C: col}
	// Authenticate the login user
	result := models.Account{}
	err = repo.C.Find(bson.M{"wallet.address": "C85ZfUB6W1KfwX3WPB9avoM4wciTKdFVbP"}).One(&result)
	fmt.Println(len(skel.ToSign))
	err = skel.Sign([]string{result.Wallet.Private})
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
