package routers

import (
	"github.com/gorilla/mux"
	"github.com/wandi34/wallets-as-a-service/backend/controllers"
)

func SetRestoreRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/restore", controllers.RestorePrivateKey).Methods("POST")
	return router
}

