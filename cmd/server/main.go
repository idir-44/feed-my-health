package main

import (
	"log"

	"github.com/idir-44/feed-my-health/pkg/database"
	"github.com/idir-44/feed-my-health/pkg/server"
)

func main() {

	srv := server.New(server.Config{Port: 8080})

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("failed to init the store: %s", err)
		return
	}

	defer db.Close()

	srv.Run()

}
