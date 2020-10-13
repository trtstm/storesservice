package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/trtstm/storesservice/cache"
	"github.com/trtstm/storesservice/metrics"
	"github.com/trtstm/storesservice/repositories"
	"github.com/trtstm/storesservice/services"
	"github.com/trtstm/storesservice/swagger/models"
	"github.com/trtstm/storesservice/swagger/restapi"
	"github.com/trtstm/storesservice/swagger/restapi/operations"
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
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	metrics := metrics.NewPrometheusMetric()

	bss := services.NewBicycleStoreService(
		repositories.NewBicycleStoreRepositoryPlaces(apiKey),
		cache.NewMemoryCache(metrics),
	)

	// Show metrics.
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":8081", nil)
	}()

	api := operations.NewStoresserviceAPI(swaggerSpec)

	api.GetBicycleStoresHandler = operations.GetBicycleStoresHandlerFunc(
		func(params operations.GetBicycleStoresParams) middleware.Responder {
			stores, err := bss.GetBicycleStoresWithinRange(lat, lon, radius)
			if err != nil {
				return operations.NewGetBicycleStoresInternalServerError().WithPayload(INTERNAL_SERVER_ERROR_MESSAGE)
			}

			searchString := ""
			if params.Search != nil {
				searchString = strings.TrimSpace(*params.Search)
			}
			stores = services.Filter(stores, services.NameFilter(searchString))

			return operations.NewGetBicycleStoresOK().WithPayload(stores)
		})

	/*if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}*/

	e := echo.New()
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)

	e.GET("/bicyclestores", echo.WrapHandler(api.Serve(nil))) // Expose our endpoint.
	e.GET("/docs", echo.WrapHandler(api.Serve(nil)))          // Expose docs
	e.GET("/swagger.json", echo.WrapHandler(api.Serve(nil)))  // And allow docs to fetch json.
	e.Logger.Fatal(e.Start(":8080"))
}
