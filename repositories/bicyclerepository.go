package repositories

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/trtstm/storesservice/models"
)

// BicycleStoreRepository returns a list of bicycle stores at a certain location.
type BicycleStoreRepository interface {
	GetBicycleStoresWithinRange(lat, lon float64, radius uint) ([]models.BicycleStore, error)
}

// BicycleStoreRepositoryPlaces implements BicycleStoreRepository for the google places api.
type BicycleStoreRepositoryPlaces struct {
	apiKey string
}

// generatePlacesUrl generates the google places url from which we can fetch the results.
func generatePlacesUrl(apiKey string, query string, lat, lon float64, radius uint) string {
	u := "https://maps.googleapis.com/maps/api/place/textsearch/json?"
	params := url.Values{}
	params.Add("key", apiKey)
	params.Add("query", query)
	params.Add("location", fmt.Sprintf("%f,%f", lat, lon))
	params.Add("radius", fmt.Sprintf("%d", radius))
	u += params.Encode()
	return u
}

// NewBicycleStoreRepositoryPlaces creates a new repository with google places as a backend.
// apiKey should be a valid google places api key.
func NewBicycleStoreRepositoryPlaces(apiKey string) *BicycleStoreRepositoryPlaces {
	return &BicycleStoreRepositoryPlaces{
		apiKey: apiKey,
	}
}

// GetBicycleStoresWithinRange gets the bicycle stores from the places api.
func (r *BicycleStoreRepositoryPlaces) GetBicycleStoresWithinRange(lat, lon float64, radius uint) ([]models.BicycleStore, error) {
	log.Printf("trying to fetch bicycle stores from places api...\n")
	placesURL := generatePlacesUrl(r.apiKey, "bicycle store", lat, lon, radius)
	resp, err := http.Get(placesURL)
	if err != nil {
		return []models.BicycleStore{}, fmt.Errorf("could not get get results from places api: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return []models.BicycleStore{}, fmt.Errorf("places returned non zero status code: %v", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []models.BicycleStore{}, fmt.Errorf("could not get read results from places api: %v", err)
	}

	results := struct {
		Results []struct {
			Name             string `json:"name"`
			FormattedAddress string `json:"formatted_address"`
		}
		Status string `json:"status"`
	}{}

	err = json.Unmarshal(body, &results)
	if err != nil {
		return []models.BicycleStore{}, fmt.Errorf("could not parse json from places api: %v", err)
	}

	if results.Status != "OK" {
		return []models.BicycleStore{}, fmt.Errorf("places returned non OK status: %v", results.Status)
	}

	bicycleStores := make([]models.BicycleStore, 0, len(results.Results)) // Pre allocate right size.
	for _, result := range results.Results {
		bicycleStores = append(bicycleStores, models.BicycleStore{Name: result.Name, Address: result.FormattedAddress})
	}

	return bicycleStores, nil
}
