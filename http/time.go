package http

import (
	"net/http"
	"time"

	"github.com/arduino-lock/golockserver"
)

// GetTime does exactly what it says
func GetTime(w http.ResponseWriter, r *http.Request, c *golockserver.Config) (int, error) {
	now := time.Now()

	_, err := w.Write([]byte(now.String()))

	return 200, err
}
