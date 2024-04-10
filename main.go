package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db, err := initStore()
	if err != nil {
		log.Fatalf("failed to init the store: %s", err)
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

func initStore() (*bun.DB, error) {
	pgConnString := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s?sslmode=disable", os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"),
		os.Getenv("PGPORT"),
		os.Getenv("PGDATABASE"),
	)

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(pgConnString)))

	db := bun.NewDB(sqldb, pgdialect.New())

	return db, nil

}
