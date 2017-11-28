package controllers

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"encoding/json"
)

func GetAddress(w http.ResponseWriter, r *http.Request) {
	// Get addressId from URL
	vars := mux.Vars(r)
	addressId := vars["id"]

	addr, err := bcy.GetAddr(addressId, nil)
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	j, err := json.Marshal(GetAddressResource{addr})
	w.Write(j)
}
