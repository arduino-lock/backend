package http

import (
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
	id := mux.Vars(r)["id"]

	card := &golockserver.Card{
		UID:     id,
		Created: time.Now(),
	}

	err := c.Services.Cards.Add(card)
	if err != nil {
		return 500, err
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

	jsonPrint(w, card)
	return http.StatusOK, nil
}

// CardGetAll fetches all cards from the database
func CardGetAll(w http.ResponseWriter, r *http.Request, c *golockserver.Config) (int, error) {
	cards, err := c.Services.Cards.GetAll()
	if err != nil {
		return http.StatusInternalServerError, err
	}

	jsonPrint(w, cards)
	return http.StatusOK, nil
}
