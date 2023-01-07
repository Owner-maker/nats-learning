package data

import (
	"errors"
	"fmt"
	"nats-learning/internal/models"
	"sync"
)

type Cache struct {
	Mu     sync.RWMutex
	Orders map[string]models.Order
}

func (cch *Cache) Put(uid string, order models.Order) {
	cch.Mu.Lock()
	defer cch.Mu.Unlock()
	cch.Orders[uid] = order
}

func (cch *Cache) Get(uid string) (models.Order, error) {
	cch.Mu.RLock()
	defer cch.Mu.RUnlock()
	var order models.Order

	if order, found := cch.Orders[uid]; found {
		return order, nil
	}

	return order, errors.New(fmt.Sprintf("order with uid:%s was not found in cache", uid))
}

func (cch *Cache) NewCache() *Cache {
	var cache Cache
	orders := make(map[string]models.Order)
	cache.Orders = orders
	return &cache
}
