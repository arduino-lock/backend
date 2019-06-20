package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arduino-lock/golockserver"
	"github.com/fatih/color"
	"github.com/gorilla/mux"
)

// Serve sets handling and starts the server
func Serve(c *golockserver.Config) {
	r := mux.NewRouter()

	r.HandleFunc("/time", Wrap(GetTime, c)).Methods("GET")
	r.HandleFunc("/validate", Wrap(CardValidate, c)).Methods("POST")

	// cards subrouter
	cards := r.PathPrefix("/cards/").Subrouter()
	cards.HandleFunc("/add/{id:[a-z]+}", Wrap(CardAdd, c)).Methods("POST")
	cards.HandleFunc("/get/all", Wrap(CardGetAll, c)).Methods("GET")
	cards.HandleFunc("/get/{id:[a-z]+}", Wrap(CardGet, c)).Methods("GET")

	fmt.Print("Listening on port ")
	color.Green(c.Port)

	log.Fatal(http.ListenAndServe(":"+c.Port, r))
}
