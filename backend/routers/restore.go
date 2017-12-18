package routers

import (
	"github.com/gorilla/mux"
	"github.com/wandi34/wallets-as-a-service/backend/controllers"
	"github.com/codegangsta/negroni"
	"github.com/wandi34/wallets-as-a-service/backend/common"
)

func SetRestoreRoutes(router *mux.Router) *mux.Router {
	transactionRouter := mux.NewRouter()
	transactionRouter.HandleFunc("/restore", controllers.CreateTransaction).Methods("POST")
	router.PathPrefix("/restore").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(transactionRouter),
	))
	return router
}

