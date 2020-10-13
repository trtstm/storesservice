package services

import (
	"errors"
	"fmt"
	"testing"

	"github.com/trtstm/storesservice/models"
)

type CacheMock struct {
	lastKey string
	useStores []models.BicycleStore
}	
func (c *CacheMock) Get(search string, getter func() ([]models.BicycleStore, error)) ([]models.BicycleStore, error) {
	c.lastKey = search

	if c.useStores != nil {
		return c.useStores, nil
	}

	return getter()
}

type RepoMock struct {
	nCalls int
	stores []models.BicycleStore
	err error
}
func (r *RepoMock) GetBicycleStoresWithinRange(lat, lon float64, radius uint) ([]models.BicycleStore, error) {
	if r.err != nil {
		return r.stores, r.err
	}
	return r.stores, nil
}

// TestBicycleStoreServiceUseOfCache tests if the service is correctly using the cache.
func TestBicycleStoreServiceUseOfCache(t *testing.T) {
	cache := &CacheMock{}
	repo := &RepoMock{}

	repo.stores = []models.BicycleStore{models.BicycleStore{Name: "Name 2", Address: "Address 2"}}

	service := NewBicycleStoreService(repo, cache)

	// Pretend cache has value.
	cache.useStores = []models.BicycleStore{models.BicycleStore{Name: "Name 1", Address: "Address 1"}}
	s, err := service.GetBicycleStoresWithinRange(1.0, 1.0, 1)
	if err != nil {
		t.Fail()
	}

	if len(s) != 1 || s[0].Name != cache.useStores[0].Name || s[0].Address != cache.useStores[0].Address {
		t.Errorf("service did not return value from cache")
	}

	if repo.nCalls > 0 {
		t.Errorf("service should have used only cache but used repo also")
	}

	// Pretend cache doesn't have value.
	cache.useStores = nil
	s, err = service.GetBicycleStoresWithinRange(2.0, 2.0, 3)
	if err != nil {
		t.Fail()
	}

	if len(s) != 1 || s[0].Name != repo.stores[0].Name || s[0].Address != repo.stores[0].Address {
		t.Errorf("service did not return value from repo")
	}

	// Check if cache contains value
	if cache.lastKey != fmt.Sprintf("%f,%f,%d", 2.0, 2.0, 3) {
		t.Errorf("service did not update cache")
	}
}

// TestBicycleStoreServiceRepoFail tests if the error propagates correctly when the repo returned an error.
func TestBicycleStoreServiceRepoFail(t *testing.T) {
	cache := &CacheMock{}
	repo := &RepoMock{}
	repo.err = errors.New("crash...")
	repo.stores = []models.BicycleStore{models.BicycleStore{Name: "Name 2", Address: "Address 2"}}

	service := NewBicycleStoreService(repo, cache)
	_, err := service.GetBicycleStoresWithinRange(10.0, 10.0, 5)
	if err == nil {
		t.Errorf("expected error from service since repo crashed")
	}
}