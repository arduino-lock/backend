package boltdb

import (
	"fmt"
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
	tx, err := s.DB.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// open bucket
	b := tx.Bucket([]byte("cards"))

	// Create a string with UNIX time value
	nowUNIX := strconv.FormatInt(card.Created.Unix(), 10)

	// Save it to the database
	b.Put([]byte(card.UID), []byte(nowUNIX))
	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

// Get fetches a card (given its unique ID) from the database
func (s *CardService) Get(uid string) (*golockserver.Card, error) {
	card := &golockserver.Card{
		UID: uid,
	}

	err := s.DB.View(func(tx *bolt.Tx) error {
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

	fmt.Println(card)

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
