package main

import "fmt"

func main() {
	_, err := loadConfig("config.json")
	if err != nil {
		panic(err)
	}

	fmt.Println("Loaded the config file.")
}
