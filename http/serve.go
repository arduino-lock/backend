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
	r.HandleFunc("/validate/{id}", Wrap(CardValidate, c)).Methods("GET", "POST")

	// database subrouter
	database := r.PathPrefix("/database").Subrouter()
	database.HandleFunc("/dump", Wrap(DatabaseDump, c)).Methods("GET")

	// cards subrouter
	cards := r.PathPrefix("/cards").Subrouter()
	cards.HandleFunc("/add", Wrap(CardAdd, c)).Methods("POST")
	cards.HandleFunc("/get/all", Wrap(CardGetAll, c)).Methods("GET")
	cards.HandleFunc("/get/{id}", Wrap(CardGet, c)).Methods("GET")
	cards.HandleFunc("/delete/{id}", Wrap(CardDelete, c)).Methods("POST")

	fmt.Print("Listening on port ")
	color.Green(c.Port)

	log.Fatal(http.ListenAndServe(":"+c.Port, r))
}
