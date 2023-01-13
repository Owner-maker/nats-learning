package cache

import (
	"fmt"
	"github.com/pkg/errors"
	"nats-learning/internal/models"
)

type OrderCache struct {
	cch *Cache
}

func NewOrderCache(cch *Cache) *OrderCache {
	return &OrderCache{cch: cch}
}

func (o *OrderCache) PutOrder(uid string, order models.Order) {
	o.cch.Mutex.Lock()
	defer o.cch.Mutex.Unlock()
	o.cch.Data[uid] = order
}

func (o *OrderCache) GetOrder(uid string) (models.Order, error) {
	o.cch.Mutex.RLock()
	defer o.cch.Mutex.RUnlock()

	if orderData, found := o.cch.Data[uid]; found {
		if value, ok := orderData.(models.Order); ok {
			return value, nil
		}
		return models.Order{}, errors.New(fmt.Sprintf("failed to convert order with uid %s to its struct", uid))
	}
	return models.Order{}, errors.New(fmt.Sprintf("order with uid %s was not found in cache", uid))
}

func (o *OrderCache) GetAllOrders() ([]models.Order, error) {
	o.cch.Mutex.RLock()
	defer o.cch.Mutex.RUnlock()
	var orders []models.Order

	if len(o.cch.Data) != 0 {
		orders = make([]models.Order, len(o.cch.Data), len(o.cch.Data))
	}

	for _, valueMap := range o.cch.Data {
		valueOrder, ok := valueMap.(models.Order)
		if !ok {
			return nil, errors.New(fmt.Sprintf("failed to convert order with uid %s to its struct", valueOrder.OrderUid))
		}
		orders = append(orders, valueOrder)
	}
	return orders, nil
}
