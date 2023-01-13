package service

import (
	"nats-learning/internal/models"
	"nats-learning/internal/repository/cache"
	"nats-learning/internal/repository/postgres"
)

type Service struct {
	OrderCache    cache.OrderCache
	OrderPostgres postgres.OrderPostgres
}

func NewService(cch cache.OrderCache, ps postgres.OrderPostgres) *Service {
	return &Service{
		OrderCache:    cch,
		OrderPostgres: ps,
	}
}

func (s *Service) GetCachedOrder(uid string) (models.Order, error) {
	return s.OrderCache.GetOrder(uid)
}

func (s *Service) GetAllCachedOrders() ([]models.Order, error) {
	return s.OrderCache.GetAllOrders()
}

func (s *Service) GetAllDbOrders() ([]models.Order, error) {
	return s.OrderPostgres.GetAll()
}

func (s *Service) PutOrdersFromDbToCache() error {
	orders, err := s.GetAllDbOrders()
	if err != nil {
		return err
	}
	for i := 0; i < len(orders); i++ {
		s.PutCachedOrder(orders[i])
	}
	return nil
}

func (s *Service) PutCachedOrder(order models.Order) {
	s.OrderCache.PutOrder(order.OrderUid, order)
}

func (s *Service) PutDbOrder(order models.Order) error {
	return s.OrderPostgres.Create(order)
}
