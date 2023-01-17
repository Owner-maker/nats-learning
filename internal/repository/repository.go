package repository

import (
	"github.com/Owner-maker/nats-learning/internal/models"
	"github.com/Owner-maker/nats-learning/internal/repository/cache"
	"github.com/Owner-maker/nats-learning/internal/repository/postgres"
	"github.com/jinzhu/gorm"
)

type OrderPostgres interface {
	Create(ord models.Order) error
	GetAll() ([]models.Order, error)
}

type OrderCache interface {
	PutOrder(uid string, order models.Order)
	GetOrder(uid string) (models.Order, error)
	GetAllOrders() ([]models.Order, error)
}

type Repository struct {
	OrderPostgres
	OrderCache
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		OrderPostgres: postgres.NewOrderPostgres(db),
		OrderCache:    cache.NewOrderCache(cache.NewCache()),
	}
}
