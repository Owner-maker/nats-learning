package postgres

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"nats-learning/internal/models"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
	SslMode  string
}

func ConnectDB(c Config) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		c.Host,
		c.Port,
		c.Username,
		c.DbName,
		c.Password,
		c.SslMode,
	))
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Order{})
	db.AutoMigrate(&models.Delivery{})
	db.AutoMigrate(&models.Payment{})
	db.AutoMigrate(&models.Item{})

	return db, nil
}
