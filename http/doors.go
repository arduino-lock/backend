package http

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/arduino-lock/golockserver"
	"github.com/gorilla/mux"
)

// DoorInstall creates a new door in the database
func DoorInstall(w http.ResponseWriter, r *http.Request, c *golockserver.Config) (int, error) {
	door := &golockserver.Door{}

	// buffer to read request body
	bodyBuffer := new(bytes.Buffer)
	bodyBuffer.ReadFrom(r.Body)

	// Parse bytes into JSON format
	err := json.Unmarshal(bodyBuffer.Bytes(), door)
	if err != nil {
		return http.StatusBadRequest, err
	}

	err = c.Services.Doors.Install(door)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

// DoorGetAll creates a new door in the database
func DoorGetAll(w http.ResponseWriter, r *http.Request, c *golockserver.Config) (int, error) {
	doors, err := c.Services.Doors.GetAll()
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return jsonPrint(w, doors)
}

// DoorGetByUID creates a new door in the database
func DoorGetByUID(w http.ResponseWriter, r *http.Request, c *golockserver.Config) (int, error) {
	uid := mux.Vars(r)["uid"]

	door, err := c.Services.Doors.GetByUID(uid)
	if err != nil {
		if err.Error() == golockserver.DoorNotFound {
			return http.StatusNotFound, nil
		}

		return http.StatusInternalServerError, err
	}

	return jsonPrint(w, door)
}

// DoorUninstall creates a new door in the database
func DoorUninstall(w http.ResponseWriter, r *http.Request, c *golockserver.Config) (int, error) {
	uid := mux.Vars(r)["uid"]

	err := c.Services.Doors.Uninstall(uid)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

// DoorAddCard adds a new card to the door given both of its UIDs
func DoorAddCard(w http.ResponseWriter, r *http.Request, c *golockserver.Config) (int, error) {
	vars := mux.Vars(r)

	doorUID := vars["doorUID"]
	cardUID := vars["cardUID"]

	err := c.Services.Doors.AddCard(doorUID, cardUID)
	if err != nil {
		if err.Error() == golockserver.DoorNotFound {
			return http.StatusNotFound, nil
		}

		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

// DoorGetAllCards fetches all door's cards
func DoorGetAllCards(w http.ResponseWriter, r *http.Request, c *golockserver.Config) (int, error) {
	vars := mux.Vars(r)

	doorUID := vars["doorUID"]

	cards, err := c.Services.Doors.GetAllCards(doorUID)
	if err != nil {
		if err.Error() == golockserver.DoorNotFound {
			return http.StatusNotFound, nil
		}

		return http.StatusInternalServerError, err
	}

	return jsonPrint(w, cards)
}

// DoorRemoveCard removes a card from a door's card list
func DoorRemoveCard(w http.ResponseWriter, r *http.Request, c *golockserver.Config) (int, error) {
	vars := mux.Vars(r)

	doorUID := vars["doorUID"]
	cardUID := vars["cardUID"]

	err := c.Services.Doors.RemoveCard(doorUID, cardUID)
	if err != nil {
		if err.Error() == golockserver.DoorNotFound {
			return http.StatusNotFound, nil
		}

		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
