package controllers

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"github.com/wandi34/wallets-as-a-service/backend/common"
)

func GetAddress(w http.ResponseWriter, r *http.Request) {
	// Get addressId from URL
	vars := mux.Vars(r)
	addressId := vars["id"]

	addr, err := common.GetAddress(addressId)
	if err != nil {
		common.DisplayAppError(w, err, "Address not available", 400)
		return
	}
	w.WriteHeader(http.StatusOK)
	j, err := json.Marshal(GetAddressResource{addr})
	w.Write(j)
}
