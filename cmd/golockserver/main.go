package main

import (
	"fmt"
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

	// Setup buckets
	db.Update(func(tx *bolt.Tx) error {
		// Create cards bucket
		_, err := tx.CreateBucketIfNotExists([]byte("cards"))
		if err != nil {
			return fmt.Errorf("Cannot create cards bucket: %s", err)
		}

		// Create doors bucket
		_, err = tx.CreateBucketIfNotExists([]byte("doors"))
		if err != nil {
			return fmt.Errorf("Cannot create doors bucket: %s", err)
		}

		return nil
	})

	c := &golockserver.Config{
		Development: cfg.Development,
		Host:        cfg.Host,
		Port:        strconv.Itoa(cfg.Port),
		Services: &golockserver.Services{
			Cards: &boltdb.CardService{
				DB: db,
			},
		},
	}

	fmt.Println("Loaded the config file.")

	h.Serve(c)
}
