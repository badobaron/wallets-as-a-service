package routers

import (
	"github.com/gorilla/mux"
	"github.com/wandi34/wallets-as-a-service/backend/controllers"
)

func SetIbanRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/iban", controllers.ConvertIban).Methods("POST")
	return router
}
