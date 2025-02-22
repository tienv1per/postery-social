package main

import (
	"log"
	"postery/internal/db"
	"postery/internal/store"
)

func main() {
	addr := "postgres://admin:adminpassword@localhost/social?sslmode=disable"
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := store.NewStorage(conn)

	db.Seed(store, conn)
}
