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
		common.DisplayAppError(w, err, "Tx Error", 400)
		return
	}
	//Sign it locally
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("accounts")
	repo := &data.AccountRepository{C: col}
	// Authenticate the login user
	result := models.Account{}
	err = repo.C.Find(bson.M{"wallet.address": sourceAddress}).One(&result)
	privateKey, _ := common.Decrypt(dataResource.Data.Password, []byte(result.Wallet.Private), result.UserId.String())

	//Sign all open transactions with private key
	var signingKeys []string
	for i := 0;i < len(skel.ToSign);i++{
		signingKeys = append(signingKeys, string(privateKey[:]))
	}
	err = skel.Sign(signingKeys)
	if err != nil {
		common.DisplayAppError(w, err, "Signing Tx Error", 400)
		return
	}
	//Send TXSkeleton
	skel, err = bcy.SendTX(skel)
	if err != nil {
		common.DisplayAppError(w, err, "Sending Tx Error", 400)
		return
	}
	fmt.Printf("%+v\n", skel)

	j, _ := json.Marshal(skel)
	w.Write(j)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

