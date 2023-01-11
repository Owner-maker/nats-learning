package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"nats-learning/internal/models"
)

var DB *gorm.DB

func DbURL(host string, port string, user string, password string, dbName string, sslMode string) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		host,
		port,
		user,
		dbName,
		password,
		sslMode,
	)
}

func ConnectDB(host string, port string, user string, password string, dbName string, sslMode string) {
	db, err := gorm.Open("postgres", DbURL(host, port, user, password, dbName, sslMode))
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&models.Order{})
	db.AutoMigrate(&models.Delivery{})
	db.AutoMigrate(&models.Payment{})
	db.AutoMigrate(&models.Item{})

	DB = db
}
