package models

import (
	"gorm.io/gorm"
)

//Users store user details
type Users struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}
