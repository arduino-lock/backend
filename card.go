package golockserver

import "time"

// Card is a single identity for each of the cards
type Card struct {
	UID     string
	Created time.Time
}

// CardService is a an interface for all the methods involving cards
type CardService interface {
	Add(c *Card) error
	GetByUID(uid string) (*Card, error)
	GetAll() (*[]Card, error)
	Delete(uid string) error
}
