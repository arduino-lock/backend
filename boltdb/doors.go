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
		door.Cards = []string{}

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
func (s *DoorService) AddCard(doorUID string, cardUID string) error {
	// fetch door
	door, err := s.GetByUID(doorUID)
	if err != nil {
		// door not found
		if err.Error() == golockserver.DoorNotFound {
			return errors.New(golockserver.DoorNotFound)
		}

		return err
	}

	// check if card exists or needs to be created
	card := &golockserver.Card{
		UID: cardUID,
	}

	if err = s.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("cards"))

		cardBuff := b.Get([]byte(card.UID))

		if len(cardBuff) > 0 {
			// card exists and is parsed into struct
			if e := json.Unmarshal(cardBuff, &card); e != nil {
				return e
			}
		} else {
			card.Created = time.Now()

			// card doesn't exist and is therefore created and saved in database
			cardBuff, e := json.Marshal(card)
			if e != nil {
				return e
			}

			// save card in database
			b.Put([]byte(card.UID), cardBuff)
		}

		// update door's cards array
		door.Cards = append(door.Cards, card.UID)

		// encode door struct
		doorBuff, e := json.Marshal(door)
		if e != nil {
			return e
		}

		// save door in database
		tx.Bucket([]byte("doors")).Put([]byte(door.UID), doorBuff)

		return nil
	}); err != nil {
		return err
	}

	return nil
}

// GetAllCards fetches all cards of the given door
func (s *DoorService) GetAllCards(doorUID string) (*[]golockserver.Card, error) {
	// fetch door by uid
	door, err := s.GetByUID(doorUID)
	if err != nil {
		return nil, err
	}

	cards := []golockserver.Card{}

	for i := 0; i < len(door.Cards); i++ {
		if e := s.DB.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("cards"))

			// fetch card bytes
			cardBytes := b.Get([]byte(door.Cards[i]))

			card := golockserver.Card{}

			// decode bytes into stuct
			if e := json.Unmarshal(cardBytes, &card); e != nil {
				return e
			}

			cards = append(cards, card)

			return nil
		}); e != nil {
			return nil, e
		}
	}

	return &cards, nil
}

// RemoveCard removes a card from the given door
func (s *DoorService) RemoveCard(doorUID string, cardUID string) error {
	if err := s.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("doors"))

		// fetch door bytes by its given uid
		doorBytes := b.Get([]byte(doorUID))

		door := golockserver.Door{}
		// decode into struct
		if e := json.Unmarshal(doorBytes, &door); e != nil {
			return e
		}

		// find card index in door's cards array
		cardIndex := golockserver.FindArrElement(door.Cards, cardUID)
		if cardIndex == -1 {
			return errors.New(golockserver.CardNotFound)
		}
		door.Cards = append(door.Cards[:cardIndex], door.Cards[cardIndex+1:]...)

		// encode struct into json again
		doorBytes, e := json.Marshal(door)
		if e != nil {
			return e
		}

		b.Put([]byte(doorUID), doorBytes)

		return nil
	}); err != nil {
		return err
	}

	return nil
}
