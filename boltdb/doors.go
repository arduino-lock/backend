package boltdb

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"time"

	"github.com/arduino-lock/golockserver"
	"github.com/boltdb/bolt"
)

// DoorService is a struct that implements DoorService and provides with all
// the functions related to dababase and data management
type DoorService struct {
	DB *bolt.DB
}

// Install creates a new door
func (s *DoorService) Install(door *golockserver.Door) error {
	if err := s.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("doors"))

		// generate sha256 hash to use as door uid
		hashBytes := sha256.Sum256([]byte(time.Now().String()))

		door.UID = hex.EncodeToString(hashBytes[:])
		door.Cards = &[]string{}

		// Create bytes buffer from door struct
		buf, e := json.Marshal(door)
		if e != nil {
			return e
		}

		e = b.Put([]byte(door.UID), buf)
		if e != nil {
			return e
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

// Uninstall deletes a door given its uid
func (s *DoorService) Uninstall(uid string) error {
	if err := s.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("doors"))

		e := b.Delete([]byte(uid))
		if e != nil {
			return e
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

// GetByUID fetches a door given it uid
func (s *DoorService) GetByUID(uid string) (*golockserver.Door, error) {
	door := &golockserver.Door{}

	if err := s.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("doors"))

		doorBuff := b.Get([]byte(uid))

		// check if door exists
		if len(doorBuff) == 0 {
			return errors.New(golockserver.DoorNotFound)
		}

		// decode buffer intro struct
		if e := json.Unmarshal(doorBuff, &door); e != nil {
			return e
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return door, nil
}

// GetAll fetches all doors
func (s *DoorService) GetAll() (*[]golockserver.Door, error) {
	doors := []golockserver.Door{}

	if err := s.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("doors"))

		if e := b.ForEach(func(key []byte, buffVal []byte) error {
			// check if door exists and buffer is not empty
			// decode and append the door found
			if len(buffVal) > 0 {
				newDoor := golockserver.Door{}

				if decodeErr := json.Unmarshal(buffVal, &newDoor); decodeErr != nil {
					return decodeErr
				}

				doors = append(doors, newDoor)
			}

			return nil
		}); e != nil {
			return e
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &doors, nil
}

// AddCard adds a new card to the door
func (s *DoorService) AddCard(d *golockserver.Door, c *golockserver.Card) error {
	return nil
}

// GetCardByUID fetches the door's card with the given UID
func (s *DoorService) GetCardByUID(d *golockserver.Door, uid string) (*golockserver.Card, error) {
	return nil, nil
}

// GetAllCards fetches all cards of the given door
func (s *DoorService) GetAllCards() (*[]golockserver.Card, error) {
	return nil, nil
}

// RemoveCard removes a card from the given door
func (s *DoorService) RemoveCard(uid string) error {
	return nil
}
