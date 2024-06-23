package controllers

import (
	"net/http"

	"github.com/idir-44/feed-my-health/internal/models"
	"github.com/labstack/echo/v4"
)

func (r controller) postUser(c echo.Context) error {
	req := models.CreateUserRequest{}
	if err := c.Bind(&req); err != nil {
		return err
	}

	user, err := r.service.CreateUser(req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, user)
}
