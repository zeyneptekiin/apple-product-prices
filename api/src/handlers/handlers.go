package handlers

import (
	"apple-product-prices/api/src/utils"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

type PriceEntry struct {
	Price *float64 `bson:"price" json:"price"`
	VAT   *float64 `bson:"vat" json:"vat"`
	Date  *string  `bson:"date" json:"date"`
}

type Product struct {
	ID          string                  `bson:"_id,omitempty" json:"id,omitempty"`
	ProductID   string                  `bson:"product_id" json:"product_id"`
	ProductName string                  `bson:"product_name" json:"product_name"`
	Price       map[string][]PriceEntry `bson:"price" json:"price"`
	Name        string                  `bson:"name" json:"name"`
	Category    string                  `bson:"category" json:"category"`
	Images      string                  `bson:"images" json:"images"`
}

func GetProductsByName(w http.ResponseWriter, r *http.Request) {
	productName := r.URL.Query().Get("product_name")

	if productName == "" {
		http.Error(w, "Missing 'product_name' query parameter", http.StatusBadRequest)
		return
	}

	collection := utils.Client.Database("apple").Collection("products")
	filter := bson.M{"product_name": productName}

	var product Product
	err := collection.FindOne(context.TODO(), filter).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "Product not found", http.StatusNotFound)
		} else {
			fmt.Printf("Query error: %v\n", err)
			http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		}
		return
	}

	filteredPrice := make(map[string][]PriceEntry)
	for key, entries := range product.Price {
		if len(entries) > 0 {
			filteredPrice[key] = entries
		}
	}
	product.Price = filteredPrice

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(product); err != nil {
		fmt.Printf("Failed to encode product to JSON: %v\n", err)
		http.Error(w, "Failed to encode data", http.StatusInternalServerError)
		return
	}
}

func GetAllProductsName(w http.ResponseWriter) {
	collection := utils.Client.Database("apple").Collection("products")

	filter := bson.M{}

	projection := bson.M{"images": 3, "name": 2, "product_name": 1, "_id": 0}

	cursor, err := collection.Find(context.TODO(), filter, options.Find().SetProjection(projection))
	if err != nil {
		fmt.Println("Query error:", err)
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	var products []struct {
		ProductName string `bson:"product_name" json:"product_name"`
		Name        string `bson:"name" json:"name"`
		Images      string `bson:"images" json:"images"`
	}

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
