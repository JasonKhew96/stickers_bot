package main

func main() {
	config, err := parseConfig()
	if err != nil {
		panic(err)
	}

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
