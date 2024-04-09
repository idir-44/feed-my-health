package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
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

func initStore() (*sql.DB, error) {
	pgConnString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", os.Getenv("PGHOST"), os.Getenv("PGPORT"),
		os.Getenv("PGDATABASE"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"))

	var (
		db  *sql.DB
		err error
	)

	db, err = sql.Open("postgres", pgConnString)
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS message (value TEXT PRIMARY KEY)"); err != nil {
		return nil, err
	}

	return db, nil

}
