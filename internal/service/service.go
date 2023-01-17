package service

import (
	"github.com/Owner-maker/nats-learning/internal/models"
	"github.com/Owner-maker/nats-learning/internal/repository"
)

type Order interface {
	GetCachedOrder(uid string) (models.Order, error)
	GetAllCachedOrders() ([]models.Order, error)
	GetAllDbOrders() ([]models.Order, error)
	PutOrdersFromDbToCache() error
	PutCachedOrder(order models.Order)
	PutDbOrder(order models.Order) error
}

type Service struct {
	repository.OrderCache
	repository.OrderPostgres
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		OrderCache:    repository.OrderCache,
		OrderPostgres: repository.OrderPostgres,
	}
}
