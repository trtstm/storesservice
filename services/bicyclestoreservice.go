package bicyclestoreservice

import (
	"github.com/trtstm/storesservice/models"
	"github.com/trtstm/storesservice/repositories"
)

// BicycleStoreService contains the logic to retrieve bicycle stores within a certain radius.
type BicycleStoreService struct {
	repo repositories.BicycleStoreRepository
}

// NewBicycleStoreService creates a new service
func NewBicycleStoreService(repo repositories.BicycleStoreRepository) *BicycleStoreService {
	return &BicycleStoreService{
		repo: repo,
	}
}

func (s *BicycleStoreService) GetBicycleStoresWithinRange(lat, lon float64, radius uint) ([]models.BicycleStore, error) {
	return s.repo.GetBicycleStoresWithinRange(lat, lon, radius)
}
