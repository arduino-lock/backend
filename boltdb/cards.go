package boltdb

import (
	"github.com/boltdb/bolt"

	"github.com/arduino-lock/golockserver"
)

// CardService is a struct that implements CardService and provies with all
// the functions related to dababase and data management
type CardService struct {
	DB *bolt.DB
}

// Add adds a new card to the database
func (s *CardService) Add(card *golockserver.Card) error {
	return nil
}

// Get fetches a card (given its unique ID) from the database
func (s *CardService) Get(uid string) (*golockserver.Card, error) {
	return nil, nil
}

// GetAll fetches all cards from the database
func (s *CardService) GetAll() (*[]golockserver.Card, error) {
	return nil, nil
}

// Delete deletes a card from the database
func (s *CardService) Delete(c *golockserver.Card) error {
	return nil
}
