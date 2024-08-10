package handlers

import (
	"apple-product-prices/api/src/utils"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

type Price struct {
	Details map[string][]string `bson:"price" json:"price"`
}

type Product struct {
	ID          string `bson:"_id,omitempty" json:"id,omitempty"`
	ProductID   string `bson:"product_id" json:"product_id"`
	ProductName string `bson:"product_name" json:"product_name"`
	Price       Price  `bson:"price" json:"price"`
	Name        string `bson:"name" json:"name"`
	Category    string `bson:"category" json:"category"`
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
		fmt.Printf("Query error: %v\n", err)
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	var products []Product
	if err := cursor.All(context.TODO(), &products); err != nil {
		fmt.Printf("Failed to decode cursor data: %v\n", err)
		http.Error(w, "Failed to decode data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(products); err != nil {
		fmt.Printf("Failed to encode products to JSON: %v\n", err)
		http.Error(w, "Failed to encode data", http.StatusInternalServerError)
		return
	}
}

func GetAllProductsName(w http.ResponseWriter, r *http.Request) {
	collection := utils.Client.Database("apple").Collection("products")

	filter := bson.M{}

	projection := bson.M{"name": 2, "product_name": 1, "_id": 0}

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
