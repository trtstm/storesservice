package cache

import (
	"log"
	"sync"

	"github.com/trtstm/storesservice/models"
)

type Cache interface {
	Get(key string, getter func() ([]models.BicycleStore, error)) ([]models.BicycleStore, error)
}

type MemoryCache struct {
	lock sync.RWMutex
	db   map[string][]models.BicycleStore
}

func NewMemoryCache() *MemoryCache {
	return &MemoryCache{
		db: map[string][]models.BicycleStore{},
	}
}

func (c *MemoryCache) Get(key string, getter func() ([]models.BicycleStore, error)) ([]models.BicycleStore, error) {
	c.lock.RLock()
	if stores, ok := c.db[key]; ok {
		log.Printf("found %s in cache\n", key)
		c.lock.RUnlock()
		return stores, nil
	}
	c.lock.RUnlock()

	log.Printf("did not find %s in cache\n", key)
	c.lock.Lock()
	stores, err := getter()
	if err != nil {
		return stores, err
	}
	c.db[key] = stores
	defer c.lock.Unlock()
	return stores, nil
}
