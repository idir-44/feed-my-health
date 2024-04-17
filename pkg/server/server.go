package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Config struct {
	Port int
}

type Server struct {
	Router *echo.Echo
	config Config
}

func (s *Server) Run() {
	s.Router.Logger.Fatal(s.Router.Start(":8080"))
}

func New(config Config) Server {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "ok",
		})
	})

	return Server{e, config}
}
