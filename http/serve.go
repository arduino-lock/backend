package http

import (
	"log"
	"net/http"

	"github.com/arduino-lock/golockserver"
	"github.com/gorilla/mux"
)

// Serve sets handling and starts the server
func Serve(c *golockserver.Config) {
	r := mux.NewRouter()

	r.HandleFunc("/time", Wrap(GetTime, c))

	log.Fatal(http.ListenAndServe(":"+c.Port, r))
}
