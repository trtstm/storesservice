package cache

import (
	"testing"

	"github.com/trtstm/storesservice/models"
)

func TestMemoryCache(t *testing.T) {
	cache := NewMemoryCache()

	storesA := []models.BicycleStore{
		{Name: "A", Address: "C"},
	}

	storesB := []models.BicycleStore{
		{Name: "D", Address: "E"},
	}

	s, err := cache.Get("a", func() ([]models.BicycleStore, error) {
		return storesA, nil
	})
	if err != nil {
		t.Errorf("expected non nil err")
	}
	if len(s) != 1 && s[0].Name != storesA[0].Name || s[0].Address != storesA[0].Address {
		t.Errorf("cache did not return correct stores")
	}

	// Test that it gets from the cache now.
	s, err = cache.Get("a", func() ([]models.BicycleStore, error) {
		t.Errorf("cache should not call getter")
		return storesA, nil
	})
	if err != nil {
		t.Errorf("expected non nil err")
	}
	if len(s) != 1 && s[0].Name != storesA[0].Name || s[0].Address != storesA[0].Address {
		t.Errorf("cache did not return correct stores")
	}

	s, err = cache.Get("b", func() ([]models.BicycleStore, error) {
		return storesB, nil
	})
	if err != nil {
		t.Errorf("expected non nil err")
	}
	if len(s) != 1 && s[0].Name != storesB[0].Name || s[0].Address != storesB[0].Address {
		t.Errorf("cache did not return correct stores")
	}

	// Test that it gets from the cache now.
	s, err = cache.Get("b", func() ([]models.BicycleStore, error) {
		t.Errorf("cache should not call getter")
		return storesB, nil
	})
	if err != nil {
		t.Errorf("expected non nil err")
	}
	if len(s) != 1 && s[0].Name != storesB[0].Name || s[0].Address != storesB[0].Address {
		t.Errorf("cache did not return correct stores")
	}
}

func BenchmarkMemoryCacheParallel(b *testing.B) {
	cache := NewMemoryCache()

	storesA := []models.BicycleStore{
		{Name: "A", Address: "C"},
	}

	b.RunParallel(func(pb *testing.PB) {
		s, err := cache.Get("a", func() ([]models.BicycleStore, error) {
			return storesA, nil
		})

		if err != nil {
			b.Errorf("expected non nil err")
		}
		if len(s) != 1 && s[0].Name != storesA[0].Name || s[0].Address != storesA[0].Address {
			b.Errorf("cache did not return correct stores")
		}
	})
}
