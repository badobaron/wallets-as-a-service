package routers

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	// Routes for the User entity
	router = SetUserRoutes(router)
	// Routes for the CryptoAccount entity
	router = SetAccountRoutes(router)
	// Routes for the Transactions entity
	router = SetTransactionsRoutes(router)
	return router
}
