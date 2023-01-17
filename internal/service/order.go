package service

import (
	"github.com/Owner-maker/nats-learning/internal/models"
)

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
