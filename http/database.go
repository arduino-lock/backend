package http

import (
	"net/http"

	"github.com/arduino-lock/golockserver"
)

// DatabaseDump calls a database dumping method
func DatabaseDump(w http.ResponseWriter, r *http.Request, c *golockserver.Config) (int, error) {
	c.Services.Database.DatabaseDump(c.Development)

	return 200, nil
}
