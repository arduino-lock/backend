package golockserver

//Door is the structure that represents every door installed
type Door struct {
	UID         string
	Cards       []string
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
	AddCard(doorUID string, cardUID string) error
	GetAllCards(doorUID string) (*[]Card, error)
	RemoveCard(doorUID string, cardUID string) error
}
