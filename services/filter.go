package services

import (
	"strings"

	"github.com/trtstm/storesservice/models"
)

// FilterFunc is a function that receives a bicycle store and returns whether to include it or not.
type FilterFunc func(models.BicycleStore) bool

// NameFilter creates a filter that filters on the name.
func NameFilter(name string) FilterFunc {
	return func(store models.BicycleStore) bool {
		return strings.Contains(strings.ToLower(store.Name), strings.ToLower(name))
	}
}

// Filter filters a list of stores with the filter function.
func Filter(stores []models.BicycleStore, filter FilterFunc) []models.BicycleStore {
	filteredStores := []models.BicycleStore{}
	for _, store := range stores {
		if filter(store) {
			filteredStores = append(filteredStores, store)
		}
	}

	return filteredStores
}
