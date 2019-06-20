package boltdb

import (
	"strconv"
	"time"

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
	// create a string with UNIX time value
	nowUNIX := strconv.FormatInt(card.Created.Unix(), 10)

	err := s.DB.Update(func(tx *bolt.Tx) error {
		// select cards bucket
		b := tx.Bucket([]byte("cards"))

		// put new data into the bucket
		err := b.Put([]byte(card.UID), []byte(nowUNIX))
		return err
	})

	if err != nil {
		return err
	}

	return nil
}

// GetByUID fetches a card (given its unique ID) from the database
func (s *CardService) GetByUID(uid string) (*golockserver.Card, error) {
	card := &golockserver.Card{
		UID: uid,
	}

	err := s.DB.View(func(tx *bolt.Tx) error {
		// select cards bucket
		b := tx.Bucket([]byte("cards"))

		// Get value from database
		createdStr := string(b.Get([]byte(uid)))

		// Convert UNIX timestamp string to int64
		var convErr error
		createdUnix, convErr := strconv.ParseInt(createdStr, 10, 64)
		if convErr != nil {
			return convErr
		}

		card.Created = time.Unix(createdUnix, 0)

		return nil
	})

	if err != nil {
		return card, err
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
			// parse unix epoch string to int64
			unixInt, err := strconv.ParseInt(string(val), 10, 64)
			if err != nil {
				return err
			}

			// create new card
			newCard := golockserver.Card{
				UID:     string(key),
				Created: time.Unix(unixInt, 0),
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
func (s *CardService) Delete(c *golockserver.Card) error {
	return nil
}
