package main

import (
	"strconv"

	"github.com/boltdb/bolt"

	"github.com/arduino-lock/golockserver"
	"github.com/arduino-lock/golockserver/boltdb"
	h "github.com/arduino-lock/golockserver/http"
)

func main() {
	cfg, err := loadConfig("config.json")
	if err != nil {
		panic(err)
	}

	// Open the database
	db, err := bolt.Open(cfg.DatabasePath, 0600, nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	c := &golockserver.Config{
		Development: cfg.Development,
		Host:        cfg.Host,
		Port:        strconv.Itoa(cfg.Port),
		Services: &golockserver.Services{
			Cards: &boltdb.CardService{
				DB: db,
			},
			Database: &boltdb.DatabaseService{
				DB: db,
			},
			Doors: &boltdb.DoorService{
				DB: db,
			},
		},
	}

	c.Services.Database.Setup(db)

	h.Serve(c)
}
