package main

import (
	"fmt"
	"strconv"

	"github.com/arduino-lock/golockserver"
	h "github.com/arduino-lock/golockserver/http"
)

func main() {
	cfg, err := loadConfig("config.json")
	if err != nil {
		panic(err)
	}

	c := &golockserver.Config{
		Development: cfg.Development,
		Host:        cfg.Host,
		Port:        strconv.Itoa(cfg.Port),
	}

	fmt.Println("Loaded the config file.")

	h.Serve(c)
}
