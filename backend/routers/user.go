package routers

import (
	"github.com/gorilla/mux"
	"github.com/wandi34/wallets-as-a-service/backend/controllers"
)

func SetUserRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/users/register", controllers.Register).Methods("POST", "OPTIONS")
	router.HandleFunc("/users/login", controllers.Login).Methods("POST", "OPTIONS")
	return router
}
