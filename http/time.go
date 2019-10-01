package http

import (
	"net/http"
	"time"

	"github.com/arduino-lock/golockserver"
	"github.com/fatih/color"
)

// GetTime does exactly what it says
func GetTime(w http.ResponseWriter, r *http.Request, c *golockserver.Config) (int, error) {
	now := time.Now()

	color.Green("%s - local time.", r.RemoteAddr)

	_, err := w.Write([]byte(now.String()))

	return 200, err
}
