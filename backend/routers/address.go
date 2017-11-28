package routers

import (
	"github.com/gorilla/mux"
	"github.com/wandi34/wallets-as-a-service/backend/controllers"
)

func SetAddressRoutes(router *mux.Router) *mux.Router {
	addressRouter := mux.NewRouter()
	addressRouter.HandleFunc("/addresses/{id}", controllers.GetAddress).Methods("GET")
	return addressRouter
}
