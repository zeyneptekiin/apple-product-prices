package main

import (
	"apple-product-prices/api/src/routes"
	"apple-product-prices/api/src/utils"
	"log"
	"net/http"

	_ "apple-product-prices/api/src/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			return
		}
		next.ServeHTTP(w, r)
	})
}

// @title Apple Product Prices API
// @version 1.0
// @description This is an API to get Apple product prices.
// @termsOfService http://swagger.io/terms/

func main() {
	utils.InitMongo()

	r := routes.SetupRoutes()

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	corsHandler := CORSMiddleware(r)

	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}
