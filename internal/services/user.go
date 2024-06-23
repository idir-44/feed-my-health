package services

import "github.com/idir-44/feed-my-health/internal/models"

func (s service) CreateUser(req models.CreateUserRequest) (models.User, error) {

	user := models.User{
		Email:    req.Email,
		Password: req.Password,
	}
	return s.repository.CreateUser(user)

}
