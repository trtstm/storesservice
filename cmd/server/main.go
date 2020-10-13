package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/trtstm/storesservice/cache"
	"github.com/trtstm/storesservice/models"
	"github.com/trtstm/storesservice/repositories"
	"github.com/trtstm/storesservice/services"
)

const INTERNAL_SERVER_ERROR_MESSAGE = "Woops something went wrong"

var apiKey string
var lat float64
var lon float64
var radius uint

func init() {
	flag.StringVar(&apiKey, "key", "", "Google places api key.")
	flag.Float64Var(&lat, "lat", 59.3317438, "Latitude. Defaults to sergels torg.")
	flag.Float64Var(&lon, "lon", 18.0647175, "Lontitude. Defaults to sergels torg.")
	flag.UintVar(&radius, "radius", 2000, "Radius in which to search.")

	flag.Parse()

	// If not set use env.
	if apiKey == "" {
		apiKey = os.Getenv("API_KEY")
	}
}

// BicycleStores is needed to marshal xml correctly.
type BicycleStores struct {
	BicycleStore []models.BicycleStore
}

func main() {
	bss := services.NewBicycleStoreService(
		repositories.NewBicycleStoreRepositoryPlaces(apiKey),
		cache.NewMemoryCache(),
	)

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.GET("/bicyclestores", func(c echo.Context) error {
		searchString := strings.TrimSpace(c.QueryParam("search"))
		format := strings.TrimSpace(c.QueryParam("format"))
		if format != "json" && format != "xml" {
			format = "json" // fallback to json.
		}

		stores, err := bss.GetBicycleStoresWithinRange(lat, lon, radius) // our rest api returns stores within a radius of 2000m around sergels torg.

		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, INTERNAL_SERVER_ERROR_MESSAGE)
		}

		stores = services.Filter(stores, services.NameFilter(searchString))

		if format == "json" {
			return c.JSON(http.StatusOK, stores)
		} else if format == "xml" {
			return c.XML(http.StatusOK, BicycleStores{BicycleStore: stores})
		}

		return c.JSON(http.StatusBadRequest, "invalid format")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
