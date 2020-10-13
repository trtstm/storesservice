package services

import (
	"testing"

	"github.com/trtstm/storesservice/models"
)

// TestFilter tests if the filtering code is working.
func TestFilter(t *testing.T) {
	stores := []models.BicycleStore{
		{Name: "Store aaaa ddd"},
		{Name: "Store bbbb"},
		{Name: "Store ddd"},
	}

	results := Filter(stores, NameFilter("aaaa"))
	if len(results) != 1 {
		t.Errorf("expected 1 result")
	}
	if results[0].Name != stores[0].Name {
		t.Errorf("filter returned wrong store")
	}

	results = Filter(stores, NameFilter("bbbb"))
	if len(results) != 1 {
		t.Errorf("expected 1 result")
	}
	if results[0].Name != stores[1].Name {
		t.Errorf("filter returned wrong store")
	}

	results = Filter(stores, NameFilter("ddd"))
	if len(results) != 2 {
		t.Errorf("expected 2 results")
	}
	if results[0].Name != stores[0].Name {
		t.Errorf("filter returned wrong store")
	}
	if results[1].Name != stores[2].Name {
		t.Errorf("filter returned wrong store")
	}
}
