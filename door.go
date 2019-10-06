package golockserver

//Door is the structure that represents every door installed
type Door struct {
	UID         string
	Cards       *[]string
	Description string
}

//DoorService is an interface for all the methods involving each door
type DoorService interface {
	// Door instance
	Install(d *Door) error
	Uninstall(uid string) error
	GetByUID(uid string) (*Door, error)
	GetAll() (*[]Door, error)

	// Door's cards
	AddCard(d *Door, c *Card) error
	GetCardByUID(d *Door, uid string) (*Card, error)
	GetAllCards() (*[]Card, error)
	RemoveCard(uid string) error
}
