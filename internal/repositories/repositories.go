package repositories

import (
	"github.com/idir-44/feed-my-health/internal/models"
	"github.com/uptrace/bun"
)

type repository struct {
	db *bun.DB
}

func NewRepository(db *bun.DB) Repository {
	return repository{db}
}

type Repository interface {
	// User
	CreateUser(req models.User) (models.User, error)
}
