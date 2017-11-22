package main

import (
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	common "github.com/wandi34/wallets-as-a-service/backend/common"
	"github.com/wandi34/wallets-as-a-service/backend/routers"
	"github.com/rs/cors"
)

//Entry point of the program
func main() {
	// allow OPTIONS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://127.0.0.1:8080", "http://localhost:8080"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"X-Auth-Key", "X-Auth-Secret", "Content-Type"},
		Debug: true,
	})

	//common.StartUp() - Replaced with init method
	// Get the mux router object
	router := routers.InitRoutes()
	// Create a negroni instance
	n := negroni.Classic()
	n.Use(c)
	n.UseHandler(router)

	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: n,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
