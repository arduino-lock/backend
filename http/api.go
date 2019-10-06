package http

import (
	"encoding/json"
	"net/http"
)

func jsonPrint(w http.ResponseWriter, d interface{}) (int, error) {
	data, err := json.MarshalIndent(d, "", "\t")
	if err != nil {
		return http.StatusInternalServerError, err
	}

	w.Write(data)
	return 0, nil
}
