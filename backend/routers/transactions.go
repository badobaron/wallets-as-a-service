package routers

import (
	"github.com/gorilla/mux"
	"github.com/wandi34/wallets-as-a-service/backend/controllers"
	"github.com/codegangsta/negroni"
	"github.com/wandi34/wallets-as-a-service/backend/common"
)

func SetTransactionsRoutes(router *mux.Router) *mux.Router {
	transactionRouter := mux.NewRouter()
	transactionRouter.HandleFunc("/transactions", controllers.CreateTransaction).Methods("POST")
	router.PathPrefix("/transactions").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(transactionRouter),
	))
	return router
}
