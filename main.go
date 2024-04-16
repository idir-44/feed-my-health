package main

import (
	"log"
	"net/http"

	"github.com/idir-44/feed-my-health/pkg/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("failed to init the store: %s", err)
		return
	}

	defer db.Close()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, message{Content: "Hello world"})
	})

	e.GET("/message", handleMessage, dumbMiddleware)

	e.Logger.Fatal(e.Start(":8080"))
}

type message struct {
	Content string `json:"content"`
}

func dumbMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		println("request to /message")
		return next(c)
	}
}
func handleMessage(c echo.Context) error {

	return c.JSON(http.StatusOK, message{Content: "you requested message route"})
}
