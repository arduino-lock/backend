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

	// doors subrouter
	doors := r.PathPrefix("/doors").Subrouter()
	doors.HandleFunc("/install", Wrap(DoorInstall, c)).Methods("POST")
	doors.HandleFunc("/get/all", Wrap(DoorGetAll, c)).Methods("GET")
	doors.HandleFunc("/get/{uid}", Wrap(DoorGetByUID, c)).Methods("GET")
	doors.HandleFunc("/uninstall/{uid}", Wrap(DoorUninstall, c)).Methods("DELETE")

	fmt.Print("Listening on port ")
	color.Green(c.Port)

	log.Fatal(http.ListenAndServe(":"+c.Port, r))
}
