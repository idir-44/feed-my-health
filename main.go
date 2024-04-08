package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

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

	return c.JSON(http.StatusOK, message{Content: "Hello world"})
}
