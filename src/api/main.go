package main

import (
	"apple-product-prices/src/api/routes"
	"apple-product-prices/src/api/utils"
	"log"
	"net/http"
)

func main() {
	utils.InitMongo()

	r := routes.SetupRoutes()

	log.Fatal(http.ListenAndServe(":8080", r))
}
