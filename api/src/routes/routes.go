package routes

import (
	"apple-product-prices/api/src/handlers"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/products", handlers.GetProductsByName).Methods("GET")
	r.HandleFunc("/products/name", handlers.GetAllProductsName).Methods("GET")
	return r
}
