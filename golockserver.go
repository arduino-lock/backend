package golockserver

// Config is a global Go LockServer config file for the multiple modules used
type Config struct {
	Development bool
	Host        string
	Port        string
	Services    *Services
}

// Services is a struct that joins all the smaller services of the app
type Services struct {
	Cards    CardService
	Database DatabaseService
	Doors    DoorService
}
