package store

import "github.com/thakkarkaushal/http-service-clean-architecture/models"

type UserStore interface {
	FindAll() ([]models.Users, error)
	FindByID(userid int) ([]models.Users, error)
}
