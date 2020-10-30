package config

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/thakkarkaushal/http-service-clean-architecture/models"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "httpservice"
	password = "http"
	dbname   = "userdetails"
)

// DB referencing the database
var DB *gorm.DB

//Connection connects the database and creates tables
func Connection() {
	dbInfo := fmt.Sprintf("host = %s port = %d user = %s password = %s dbname = %s", host, port, user, password, dbname)
	db, err := gorm.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal(err)
	}
	db.SingularTable(true)
	// db.DropTableIfExists(&models.Users{})
	db.AutoMigrate(&models.Users{})
	DB = db
	fmt.Println("Connect success")
}
