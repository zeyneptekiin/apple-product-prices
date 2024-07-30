package handlers

import (
	"apple-product-prices/src/api/utils"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

type Product struct {
	Name    string `json:"name"`
	Price   int    `json:"price"`
	Country string `json:"country"`
}

func GetProductsByName(w http.ResponseWriter, r *http.Request) {
	productName := r.URL.Query().Get("product_name")

	if productName == "" {
		http.Error(w, "Missing 'product_name' query parameter", http.StatusBadRequest)
		return
	}

	collection := utils.Client.Database("apple").Collection("products")
	filter := bson.M{"product_name": productName}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("Query error:", err)
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	var products []Product
	if err = cursor.All(context.TODO(), &products); err != nil {
		fmt.Println("Failed to decode cursor data:", err)
		http.Error(w, "Failed to decode data", http.StatusInternalServerError)
		return
	}

	fmt.Println("Decoded Products:", products)

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(products); err != nil {
		fmt.Println("Failed to encode products to JSON:", err)
		http.Error(w, "Failed to encode data", http.StatusInternalServerError)
	}
}
