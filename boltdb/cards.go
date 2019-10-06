package boltdb

import (
	"encoding/json"
	"errors"

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
	err := s.DB.Update(func(tx *bolt.Tx) error {
		// select cards bucket
		b := tx.Bucket([]byte("cards"))

		// Encode Card struct to json
		cardJSON, e := json.Marshal(card)
		if e != nil {
			return e
		}

		// put new data into the bucket
		err := b.Put([]byte(card.UID), cardJSON)
		return err
	})

	if err != nil {
		return err
	}

	return nil
}

// GetByUID fetches a card (given its unique ID) from the database
func (s *CardService) GetByUID(uid string) (*golockserver.Card, error) {
	var card *golockserver.Card

	if err := s.DB.View(func(tx *bolt.Tx) error {
		// select cards bucket
		b := tx.Bucket([]byte("cards"))

		// Get value from database
		cardBytes := b.Get([]byte(uid))

		// Check if card exists
		if len(cardBytes) == 0 {
			return errors.New(golockserver.CardNotFound)
		}

		// Decode it
		if e := json.Unmarshal(cardBytes, &card); e != nil {
			return e
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return card, nil
}

// GetAll fetches all cards from the database
func (s *CardService) GetAll() (*[]golockserver.Card, error) {
	cards := []golockserver.Card{}

	if err := s.DB.View(func(tx *bolt.Tx) error {
		// select cards bucket
		b := tx.Bucket([]byte("cards"))

		if err := b.ForEach(func(key []byte, val []byte) error {
			var newCard golockserver.Card

			e := json.Unmarshal(val, &newCard)
			if e != nil {
				return e
			}

			// append new card to cards list
			cards = append(cards, newCard)

			return nil
		}); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return &cards, err
	}

	return &cards, nil
}

// Delete deletes a card from the database
func (s *CardService) Delete(uid string) error {
	if err := s.DB.Update(func(tx *bolt.Tx) error {
		if e := tx.Bucket([]byte("cards")).Delete([]byte(uid)); e != nil {
			return e
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
