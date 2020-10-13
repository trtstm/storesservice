package services

import (
	"fmt"

	"github.com/trtstm/storesservice/cache"
	"github.com/trtstm/storesservice/metrics"
	"github.com/trtstm/storesservice/repositories"
	"github.com/trtstm/storesservice/swagger/models"
)

// BicycleStoreService contains the logic to retrieve bicycle stores within a certain radius.
type BicycleStoreService struct {
	repo    repositories.BicycleStoreRepository
	cache   cache.Cache
	metrics metrics.Metrics
}

// NewBicycleStoreService creates a new service.
// Requires a repository and cache to work.
func NewBicycleStoreService(repo repositories.BicycleStoreRepository, cache cache.Cache, metrics metrics.Metrics) *BicycleStoreService {
	return &BicycleStoreService{
		repo:    repo,
		cache:   cache,
		metrics: metrics,
	}
}

// GetBicycleStoresWithinRange uses the repository and caching to retrieve a list of bicycle stores.
func (s *BicycleStoreService) GetBicycleStoresWithinRange(lat, lon float64, radius uint) ([]*models.BicycleStore, error) {
	if s.metrics != nil {
		s.metrics.IncRequests()
	}

	// We cache on the lat,lon,radius pairs.
	return s.cache.Get(fmt.Sprintf("%f,%f,%d", lat, lon, radius), func() ([]*models.BicycleStore, error) {
		return s.repo.GetBicycleStoresWithinRange(lat, lon, radius)
	})
}
