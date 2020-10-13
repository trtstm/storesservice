package cache

import (
	"log"
	"sync"

	"github.com/trtstm/storesservice/metrics"
	"github.com/trtstm/storesservice/swagger/models"
)

type Cache interface {
	Get(key string, getter func() ([]*models.BicycleStore, error)) ([]*models.BicycleStore, error)
}

type MemoryCache struct {
	lock    sync.RWMutex
	db      map[string][]*models.BicycleStore
	metrics metrics.Metrics
}

func NewMemoryCache(metrics metrics.Metrics) *MemoryCache {
	return &MemoryCache{
		db:      map[string][]*models.BicycleStore{},
		metrics: metrics,
	}
}

// Get gets the key from cache or calls the getter function if not found.
func (c *MemoryCache) Get(key string, getter func() ([]*models.BicycleStore, error)) ([]*models.BicycleStore, error) {
	c.lock.RLock()
	if stores, ok := c.db[key]; ok {
		log.Printf("found %s in cache\n", key)
		if c.metrics != nil {
			c.metrics.IncCacheHits()
		}
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
