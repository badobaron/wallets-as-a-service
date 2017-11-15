package routers

import (
	"github.com/gorilla/mux"
	"github.com/wandi34/wallets-as-a-service/backend/controllers"
	"github.com/codegangsta/negroni"
	"github.com/wandi34/wallets-as-a-service/backend/common"
)

func SetAccountRoutes(router *mux.Router) *mux.Router {
	accountRouter := mux.NewRouter()
	accountRouter.HandleFunc("/users/{id}/accounts", controllers.CreateAccount).Methods("POST")
	accountRouter.HandleFunc("/users/{id}/accounts", controllers.GetAccounts).Methods("GET")
	router.PathPrefix("/users/{id}/accounts").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(accountRouter),
	))
	return router
}
