package routes

import (
	"apple-product-prices/api/src/handlers"
	"github.com/gorilla/mux"

	_ "apple-product-prices/api/src/docs"
)

// SetupRoutes initializes all API routes and returns the router.
// @title Apple Product Prices API Routes
// @version 1.0
// @BasePath /
// @tag.name Products
// @tag.description Operations related to Apple products
// @tag.name Exchange
// @tag.description Operations related to exchange rates

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// @Summary Get products by name
	// @Description Retrieves a list of Apple products filtered by name.
	// @Tags Products
	// @Accept  json
	// @Produce  json
	// @Success 200 {array} models.Product
	// @Failure 400 {object} models.ErrorResponse
	// @Router /products [get]
	r.HandleFunc("/products", handlers.GetProductsByName).Methods("GET")

	// @Summary Get all product names
	// @Description Retrieves a list of all Apple product names.
	// @Tags Products
	// @Accept  json
	// @Produce  json
	// @Success 200 {array} string
	// @Failure 400 {object} models.ErrorResponse
	// @Router /products/name [get]
	r.HandleFunc("/products/name", handlers.GetAllProductsName).Methods("GET")

	// @Summary Get exchange rates
	// @Description Retrieves the latest exchange rates.
	// @Tags Exchange
	// @Accept  json
	// @Produce  json
	// @Success 200 {object} models.ExchangeRates
	// @Failure 400 {object} models.ErrorResponse
	// @Router /exchange [get]
	r.HandleFunc("/exchange", handlers.GetExchangeRates).Methods("GET")

	return r
}
