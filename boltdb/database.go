package boltdb

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/arduino-lock/golockserver"
	"github.com/boltdb/bolt"
	"github.com/fatih/color"
)

// DatabaseService is a struct which implements DatabaseService interface and thus,
// implements all of its methods which perform database management related actions
type DatabaseService struct {
	DB *bolt.DB
}

// All constants
const (
	CardNotFound = "CardNotFound"
)

// Setup is a function that generates all buckets inside the db
func (s *DatabaseService) Setup(db *bolt.DB) error {
	BUCKETS := []string{"cards", "doors"}

	for i := 0; i < len(BUCKETS); i++ {
		// Create each bucket if it doesn't exist yet
		db.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists([]byte(BUCKETS[i]))
			if err != nil {
				return fmt.Errorf("Cannot create cards bucket: %s", err)
			}

			return nil
		})
	}

	return nil
}

// DatabaseDump is a function that prints all data in the database
func (s *DatabaseService) DatabaseDump(development bool) error {
	BUCKETS := []string{"cards"}

	if development {
		for i := 0; i < len(BUCKETS); i++ {
			currentBucket := BUCKETS[i]

			if err := s.DB.View(func(tx *bolt.Tx) error {
				// select cards bucket
				b := tx.Bucket([]byte(currentBucket))

				fmt.Printf("Bucket: %s\n", currentBucket)
				items := 0

				if err := b.ForEach(func(key []byte, val []byte) error {
					items++

					var card *golockserver.Card

					// decode card into struct instance
					e := json.Unmarshal(val, card)
					if e != nil {
						return e
					}
					fmt.Println(card)

					if currentBucket == "cards" {
						fmt.Printf("Card UID: ")
						color.Green(string(key))
					}

					return nil
				}); err != nil {
					return err
				}

				if items == 0 {
					color.Yellow("No records found")
				}

				return nil
			}); err != nil {
				log.Fatal(err)
			}
		}
	} else {
		color.Red("Cannot dump database while server is not running on development mode")
	}

	return nil
}
