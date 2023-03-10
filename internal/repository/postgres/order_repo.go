package postgres

import (
	"github.com/Owner-maker/nats-learning/internal/models"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type OrderPostgresRepo struct {
	db *gorm.DB
}

func NewOrderPostgres(db *gorm.DB) *OrderPostgresRepo {
	return &OrderPostgresRepo{db: db}
}

func (o *OrderPostgresRepo) Create(ord models.Order) error {
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

func (o *OrderPostgresRepo) GetAll() ([]models.Order, error) {
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

func (o *OrderPostgresRepo) Get(uid string) (models.Order, error) {
	var order models.Order
	err := o.db.Transaction(func(tx *gorm.DB) error {
		if err := o.db.
			Model(&models.Order{}).
			Preload("Delivery").
			Preload("Payment").
			Preload("Items").
			Where("order_uid = ?", uid).
			First(&order).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		logrus.Error(err)
		return models.Order{}, err
	}
	return order, nil
}
