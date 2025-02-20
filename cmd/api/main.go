package main

import (
	"log"
	"postery/internal/env"
	"postery/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	store := store.NewStorage(nil)
	
	app := &application{
		store:  store,
		config: cfg,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
