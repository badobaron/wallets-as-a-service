package main

import (
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	common "github.com/wandi34/wallets-as-a-service/backend/common"
	"github.com/wandi34/wallets-as-a-service/backend/routers"
)

//Entry point of the program
func main() {
	//common.StartUp() - Replaced with init method
	// Get the mux router object
	router := routers.InitRoutes()
	// Create a negroni instance
	n := negroni.Classic()
	n.UseHandler(router)

	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: n,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
