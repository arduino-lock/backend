package boltdb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/fatih/color"
)

// DatabaseService is a struct which implements DatabaseService interface and thus,
// implements all of its methods which perform database management related actions
type DatabaseService struct {
	DB *bolt.DB
}

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

// DatabasePrint is a function that prints all data in the database
func (s *DatabaseService) DatabasePrint(development bool) error {
	BUCKETS := []string{"cards", "doors"}

	if development {
		for i := 0; i < len(BUCKETS); i++ {
			currentBucket := BUCKETS[i]

			if err := s.DB.View(func(tx *bolt.Tx) error {
				// select cards bucket
				b := tx.Bucket([]byte(currentBucket))

				fmt.Printf("Bucket: %s\n", currentBucket)
				items := 0

				if err := b.ForEach(func(key []byte, val []byte) error {
					// parse unix epoch string to int64
					_, err := strconv.ParseInt(string(val), 10, 64)
					if err != nil {
						return err
					}

					items++

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
