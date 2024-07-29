package main

import (
	"apple-product-prices/src/api/routes"
	"apple-product-prices/src/api/utils"
	"log"
	"net/http"
)

type product struct {
	Id struct {
		Oid string `json:"$oid"`
	} `json:"_id"`
	ProductId string `json:"product_id"`
	Name      string `json:"name"`
	Price     int    `json:"price"`
	Country   string `json:"country"`
	Category  string `json:"category"`
}

func main() {
	utils.InitMongo()

	r := routes.SetupRoutes()

	log.Fatal(http.ListenAndServe(":8080", r))
}
