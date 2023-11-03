package main

import "fmt"

func main() {
	config, err := parseConfig()
	fmt.Println(config, err)

	db, err := NewDatabase()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = NewBot(config, db)
	if err != nil {
		panic(err)
	}

}
