package main

import (
	"fmt"
	"log"
	"net/http"
)

//Handler logic into a Closure
func messageHandler(message string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, message)
	})
	}
	

func main() {
	mux := http.NewServeMux()

	mux.Handle("/welcome", messageHandler("Welcome to this app"))
	mux.Handle("/", messageHandler("This is the mainmage, navigate to welcome"))

	log.Println("Listening")

	http.ListenAndServe(":8080", mux)

}