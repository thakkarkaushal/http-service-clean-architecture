package services

import (
	"github.com/thakkarkaushal/http-service-clean-architecture/models"
	"github.com/thakkarkaushal/http-service-clean-architecture/store"
)

type service struct {
	repo store.UserStore
}

func New(repos store.UserStore) UserService {
	return service{repo: repos}
}

func (s service) FindAll() ([]models.Users, error) {
	return s.repo.FindAll()
}

func (s service) FindByID(userid int) ([]models.Users, error) {
	return s.repo.FindByID(userid)
}
