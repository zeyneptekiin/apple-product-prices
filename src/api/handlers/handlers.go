package handlers

import (
	"apple-product-prices/src/api/utils"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

// Product yapısını tanımla
type Product struct {
	ID struct {
		Oid string `json:"$oid"`
	} `json:"_id"`
	ProductId string `json:"product_id"`
	Name      string `json:"name"`
	Price     int    `json:"price"`
	Country   string `json:"country"`
	Category  string `json:"category"`
}

// GetProductsByName, belirli bir ürün adı için ürünleri getirir
func GetProductsByName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name") // URL'den ürün adı al

	if name == "" {
		http.Error(w, "Missing 'name' query parameter", http.StatusBadRequest)
		return
	}

	collection := utils.Client.Database("your_database_name").Collection("your_collection_name")
	filter := bson.M{"name": name}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	var products []Product
	if err = cursor.All(context.TODO(), &products); err != nil {
		http.Error(w, "Failed to decode data", http.StatusInternalServerError)
		return
	}

	// JSON formatında yanıt ver
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
