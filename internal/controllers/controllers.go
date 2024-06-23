package controllers

import (
	"github.com/idir-44/feed-my-health/internal/services"
	"github.com/idir-44/feed-my-health/pkg/server"
)

type controller struct {
	service services.Service
}

func RegisterHandlers(routerGroup *server.Router, service services.Service) {
	c := controller{service}

	routerGroup.POST("/users", c.postUser)
}
