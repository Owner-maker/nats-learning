package postgres

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"nats-learning/internal/models"
)

type OrderPostgres struct {
	db *gorm.DB
}

func NewOrderPostgres(db *gorm.DB) *OrderPostgres {
	return &OrderPostgres{db: db}
}

func (o *OrderPostgres) Create(ord models.Order) error {
	err := o.db.Transaction(func(tx *gorm.DB) error {
		if err := o.db.Create(&ord).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		logrus.Print(err)
		return err
	}
	return nil
}

func (o *OrderPostgres) GetAll() ([]models.Order, error) {
	var orders []models.Order
	err := o.db.Transaction(func(tx *gorm.DB) error {
		if err := o.db.
			Model(&models.Order{}).
			Preload("Delivery").
			Preload("Payment").
			Preload("Items").
			Find(&orders).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		logrus.Fatal(err)
		return nil, err
	}
	return orders, nil
}
