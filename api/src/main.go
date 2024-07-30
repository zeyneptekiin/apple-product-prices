package main

import (
	"apple-product-prices/api/src/routes"
	"apple-product-prices/api/src/utils"
	"log"
	"net/http"
)

func main() {
	utils.InitMongo()

	r := routes.SetupRoutes()

	log.Fatal(http.ListenAndServe(":8080", r))
}
