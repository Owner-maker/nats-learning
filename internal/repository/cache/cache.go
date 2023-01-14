package cache

import (
	"sync"
)

type Cache struct {
	Mutex sync.RWMutex
	Data  map[string]interface{}
}

func NewCache() *Cache {
	var cache Cache
	cache.Data = make(map[string]interface{}, 0)
	return &cache
}
