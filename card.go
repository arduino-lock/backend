package golockserver

// Card is a single identity for each of the cards
type Card struct {
	UID              string
	CreationUNIXTime int
}

// CardService is a an interface for all the methods involving cards
type CardService interface {
	Add(c *Card) error
	Delete(uid string) error
}
