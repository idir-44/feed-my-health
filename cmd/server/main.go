package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/idir-44/feed-my-health/internal/controllers"
	"github.com/idir-44/feed-my-health/internal/repositories"
	"github.com/idir-44/feed-my-health/internal/services"
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

	repository := repositories.NewRepository(db)
	service := services.NewService(repository)

	v1 := srv.NewGroup("/v1")

	controllers.RegisterHandlers(v1, service)

	data, err := json.MarshalIndent(srv.Router.Routes(), "", "  ")
	if err != nil {
		fmt.Printf("failed to marshal routes: %s", err)
	}

	fmt.Println(string(data))

	srv.Run()

}
