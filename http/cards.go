package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/arduino-lock/golockserver"
)

// CardValidate is a card validator
func CardValidate(w http.ResponseWriter, r *http.Request, c *golockserver.Config) (int, error) {
	vars := mux.Vars(r)

	c.Services.Cards.Get(vars["id"])
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

	return 200, nil
}

// CardGet fetches a card from the database given its UID
func CardGet(w http.ResponseWriter, r *http.Request, c *golockserver.Config) (int, error) {
	id := mux.Vars(r)["id"]
	fmt.Println("E que crl")

	card, err := c.Services.Cards.Get(id)
	if err != nil {
		return 500, err
	}

	fmt.Println(card)
	return 200, nil
}
