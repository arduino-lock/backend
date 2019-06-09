package main

import (
	"fmt"
	"strconv"

	"github.com/boltdb/bolt"

	"github.com/arduino-lock/golockserver"
	h "github.com/arduino-lock/golockserver/http"
)

func main() {
	cfg, err := loadConfig("config.json")
	if err != nil {
		panic(err)
	}

	db, err := bolt.Open(cfg.DatabasePath, 0600, nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	
	
	c := &golockserver.Config{
		Development: cfg.Development,
		Host:        cfg.Host,
		Port:        strconv.Itoa(cfg.Port),
	}

	fmt.Println("Loaded the config file.")

	h.Serve(c)
}
