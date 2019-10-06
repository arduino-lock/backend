package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/fatih/color"
	"github.com/gorilla/mux"

	"github.com/arduino-lock/golockserver"
)

// CardValidate is a card validator
func CardValidate(w http.ResponseWriter, r *http.Request, c *golockserver.Config) (int, error) {
	vars := mux.Vars(r)

	fmt.Printf("%s - ", r.RemoteAddr)
	color.Yellow("card validation (UID: %s)", vars["id"])

	_, err := c.Services.Cards.GetByUID(vars["id"])
	if err != nil {
		fmt.Printf("%s - ", r.RemoteAddr)
		color.Red("card is not authorized (UID: %s)", vars["id"])
		return 404, nil
	}

	fmt.Printf("%s - ", r.RemoteAddr)
	color.Green("card is authorized (UID: %s)", vars["id"])
	return 200, nil
}

// CardAdd creates a new card in the database
func CardAdd(w http.ResponseWriter, r *http.Request, c *golockserver.Config) (int, error) {
	newCard := &golockserver.Card{
		Created: time.Now(),
	}

	// buffer to read request body
	bodyBuffer := new(bytes.Buffer)
	bodyBuffer.ReadFrom(r.Body)

	// Parse bytes into JSON format
	err := json.Unmarshal(bodyBuffer.Bytes(), newCard)
	if err != nil {
		return http.StatusBadRequest, err
	}

	// Check if a card with the same ID and the new one exists
	_, err = c.Services.Cards.GetByUID(newCard.UID)
	if err == nil {
		return http.StatusConflict, nil
	}

	err = c.Services.Cards.Add(newCard)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

// CardGet fetches a card from the database given its UID
func CardGet(w http.ResponseWriter, r *http.Request, c *golockserver.Config) (int, error) {
	id := mux.Vars(r)["id"]

	card, err := c.Services.Cards.GetByUID(id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return jsonPrint(w, card)
}

// CardGetAll fetches all cards from the database
func CardGetAll(w http.ResponseWriter, r *http.Request, c *golockserver.Config) (int, error) {
	cards, err := c.Services.Cards.GetAll()
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return jsonPrint(w, cards)
}

// CardDelete deletes a card given its id
func CardDelete(w http.ResponseWriter, r *http.Request, c *golockserver.Config) (int, error) {
	id := mux.Vars(r)["id"]

	err := c.Services.Cards.Delete(id)
	if err != nil {

	}

	return http.StatusOK, nil
}
