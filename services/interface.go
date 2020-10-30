package services

import "github.com/thakkarkaushal/http-service-clean-architecture/models"

type UserService interface {
	FindAll() ([]models.Users, error)
	FindByID(userid int) ([]models.Users, error)
}
