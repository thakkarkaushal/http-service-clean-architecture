package store

import (
	"github.com/thakkarkaushal/http-service-clean-architecture/config"
	"github.com/thakkarkaushal/http-service-clean-architecture/models"
)

type userStoreGORM struct{}

func New() UserStore {
	return userStoreGORM{}
}

func (r userStoreGORM) FindAll() ([]models.Users, error) {
	var userData []models.Users
	err := config.DB.Find(&userData).Error
	if err != nil {
		return nil, err
	}

	return userData, nil
}

func (r userStoreGORM) FindByID(userid int) ([]models.Users, error) {
	var userData []models.Users
	err := config.DB.Where("id = ?", userid).Find(&userData).Error
	if err != nil {
		return nil, err
	}
	return userData, nil
}
