package http

import (
	"encoding/json"
	"net/http"

	"github.com/arduino-lock/golockserver"
	"github.com/fatih/color"
)

type response struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
}

// LockHandler is a custom HTTP handler to include the program config
type LockHandler func(w http.ResponseWriter, r *http.Request, c *golockserver.Config) (int, error)

// Wrap function wraps a standard HTTP function handler with the config
func Wrap(h LockHandler, c *golockserver.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			code int
			err  error
		)

		// leave the data treatment and the response final handling
		// to the end of the wrapping function
		defer func() {
			// whenever a function writes anything into the response
			// don't treat the response body as usual
			if code == 0 && err == nil {
				return
			}

			// most common case is when the functions executed don't
			// write on the response body - that takes place here, after
			// everything has been done
			res := &response{
				Code: code,
			}

			if code != 0 {
				w.WriteHeader(code)
			}

			if err != nil {
				color.Red(err.Error())
			}

			// Write JSON into response body
			data, e := json.MarshalIndent(res, "", "\t")
			if e != nil {
				color.Red(e.Error())
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)

			return
		}()

		code, err = h(w, r, c)
	}
}
