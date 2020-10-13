package cache

import (
	"testing"

	"github.com/trtstm/storesservice/swagger/models"
)

func newStr(str string) *string {
	return &str
}

type MetricsMock struct {
	requests  int
	cacheHits int
}

func (m *MetricsMock) IncRequests() {
	m.requests++
}
func (m *MetricsMock) IncCacheHits() {
	m.cacheHits++
}

// TestMemoryCache tests in memory cache and if metrics are working on it.
func TestMemoryCache(t *testing.T) {
	metrics := &MetricsMock{}
	cache := NewMemoryCache(metrics)

	storesA := []*models.BicycleStore{
		&models.BicycleStore{Name: newStr("A"), Address: newStr("B")},
	}

	storesB := []*models.BicycleStore{
		&models.BicycleStore{Name: newStr("D"), Address: newStr("E")},
	}

	s, err := cache.Get("a", func() ([]*models.BicycleStore, error) {
		return storesA, nil
	})
	if err != nil {
		t.Errorf("expected non nil err")
	}
	if len(s) != 1 && s[0].Name != storesA[0].Name || s[0].Address != storesA[0].Address {
		t.Errorf("cache did not return correct stores")
	}

	// Test that it gets from the cache now.
	s, err = cache.Get("a", func() ([]*models.BicycleStore, error) {
		t.Errorf("cache should not call getter")
		return storesA, nil
	})
	if err != nil {
		t.Errorf("expected non nil err")
	}
	if len(s) != 1 && s[0].Name != storesA[0].Name || s[0].Address != storesA[0].Address {
		t.Errorf("cache did not return correct stores")
	}

	s, err = cache.Get("b", func() ([]*models.BicycleStore, error) {
		return storesB, nil
	})
	if err != nil {
		t.Errorf("expected non nil err")
	}
	if len(s) != 1 && s[0].Name != storesB[0].Name || s[0].Address != storesB[0].Address {
		t.Errorf("cache did not return correct stores")
	}

	// Test that it gets from the cache now.
	s, err = cache.Get("b", func() ([]*models.BicycleStore, error) {
		t.Errorf("cache should not call getter")
		return storesB, nil
	})
	if err != nil {
		t.Errorf("expected non nil err")
	}
	if len(s) != 1 && s[0].Name != storesB[0].Name || s[0].Address != storesB[0].Address {
		t.Errorf("cache did not return correct stores")
	}

	if metrics.cacheHits != 2 {
		t.Errorf("expected 2 cache hits")
	}
}

func BenchmarkMemoryCacheParallel(b *testing.B) {
	cache := NewMemoryCache(nil)

	storesA := []*models.BicycleStore{
		&models.BicycleStore{Name: newStr("A"), Address: newStr("C")},
	}

	b.RunParallel(func(pb *testing.PB) {
		s, err := cache.Get("a", func() ([]*models.BicycleStore, error) {
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
