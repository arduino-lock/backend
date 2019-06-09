package http

import (
	"net/http"

	"github.com/arduino-lock/golockserver"
)

// LockHandler is a custom HTTP handler to include the program config
type LockHandler func(w http.ResponseWriter, r *http.Request, c *golockserver.Config) (int, error)

// Wrap function wraps a standard HTTP function handler with the config
func Wrap(h LockHandler, c *golockserver.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code, err := h(w, r, c)
		if err != nil {
			w.WriteHeader(code)
		}
	}
}
