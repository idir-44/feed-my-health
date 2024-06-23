package services

import (
	"github.com/idir-44/feed-my-health/internal/models"
	"github.com/idir-44/feed-my-health/internal/repositories"
)

type service struct {
	repository repositories.Repository
}

func NewService(repository repositories.Repository) Service {
	return service{repository}
}

type Service interface {
	// User
	CreateUser(req models.CreateUserRequest) (models.User, error)
}
