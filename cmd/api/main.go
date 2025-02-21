package main

import (
	"database/sql"
	"log"
	"postery/internal/db"
	"postery/internal/env"
	"postery/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         "postgres://admin:adminpassword@localhost/social?sslmode=disable",
			maxOpenConns: 30,
			maxIdleConns: 30,
			maxIdleTime:  "15m",
		},
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		log.Panic(err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Panic("Error when closing DB: ", err)
		}
	}(db)

	log.Println("DB connection pool established")

	appStore := store.NewStorage(db)

	app := &application{
		store:  appStore,
		config: cfg,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
