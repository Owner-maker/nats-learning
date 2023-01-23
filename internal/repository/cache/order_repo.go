package cache

import (
	"fmt"
	"github.com/Owner-maker/nats-learning/internal/models"
	"github.com/pkg/errors"
	"net/http"
)

type OrderCacheRepo struct {
	cch *Cache
}

func NewOrderCache(cch *Cache) *OrderCacheRepo {
	return &OrderCacheRepo{cch: cch}
}

func (o *OrderCacheRepo) PutOrder(uid string, order models.Order) {
	o.cch.Mutex.Lock()
	defer o.cch.Mutex.Unlock()
	o.cch.Data[uid] = order
}

func (o *OrderCacheRepo) GetOrder(uid string) (models.Order, error) {
	o.cch.Mutex.RLock()
	defer o.cch.Mutex.RUnlock()

	if orderData, found := o.cch.Data[uid]; found {
		if value, ok := orderData.(models.Order); ok {
			return value, nil
		}
		return models.Order{},
			NewErrorHandler(
				errors.New(fmt.Sprintf("failed to convert order with uid %s to its struct", uid)),
				http.StatusInternalServerError)
	}
	return models.Order{},
		NewErrorHandler(
			errors.New(fmt.Sprintf("order with uid %s was not found in cache", uid)),
			http.StatusBadRequest)
}

func (o *OrderCacheRepo) GetAllOrders() ([]models.Order, error) {
	o.cch.Mutex.RLock()
	defer o.cch.Mutex.RUnlock()

	if len(o.cch.Data) == 0 {
		return []models.Order{}, nil
	}
	orders := make([]models.Order, len(o.cch.Data), len(o.cch.Data))

	i := 0
	for _, valueMap := range o.cch.Data {
		valueOrder, ok := valueMap.(models.Order)
		if !ok {
			return nil,
				NewErrorHandler(
					errors.New(fmt.Sprintf("failed to convert order with uid %s to its struct", valueOrder.OrderUid)),
					http.StatusInternalServerError)
		}
		orders[i] = valueOrder
		i++
	}
	return orders, nil
}
