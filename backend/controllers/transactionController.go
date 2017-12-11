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
	err = repo.C.Find(bson.M{"wallet.address": sourceAddress}).One(&result)
	fmt.Println(len(skel.ToSign))
	// Decrypt private key
	privateKey, _ := common.Decrypt(common.GetMd5Hash("1234"), []byte(result.Wallet.Private))

	//Sign all open transactions with private key
	var signingKeys []string
	for i := 0;i < len(skel.ToSign);i++{
		signingKeys = append(signingKeys, string(privateKey[:]))
	}
	err = skel.Sign(signingKeys)
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

