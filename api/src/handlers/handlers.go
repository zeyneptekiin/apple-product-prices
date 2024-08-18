package handlers

import (
	"apple-product-prices/api/src/utils"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
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

type ExchangeRatesResponse struct {
	Result             string             `json:"result"`
	TimeLastUpdateUnix int64              `json:"time_last_update_unix"`
	TimeLastUpdateUTC  string             `json:"time_last_update_utc"`
	BaseCode           string             `json:"base_code"`
	ConversionRates    map[string]float64 `json:"conversion_rates"`
}

// GetProductsByName handles the request to retrieve a product by its name.
// @Summary Get a product by name
// @Description Retrieves a single Apple product by its name from the database.
// @Tags Products
// @Accept  json
// @Produce  json
// @Param product_name query string true "Name of the product"
// @Success 200 {object} Product
// @Failure 400 {object} map[string]string "Missing 'product_name' query parameter"
// @Failure 404 {object} map[string]string "Product not found"
// @Failure 500 {object} map[string]string "Failed to fetch data"
// @Router /products [get]
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

// GetAllProductsName handles the request to retrieve all product names.
// @Summary Get all product names
// @Description Retrieves a list of all Apple product names, including images.
// @Tags Products
// @Accept  json
// @Produce  json
// @Success 200 {array} map[string]string
// @Failure 500 {object} map[string]string "Failed to fetch data"
// @Router /products/name [get]
func GetAllProductsName(w http.ResponseWriter, r *http.Request) {
	collection := utils.Client.Database("apple").Collection("products")

	filter := bson.M{}

	projection := bson.M{"images": 3, "name": 2, "product_name": 1, "_id": 0}

	cursor, err := collection.Find(context.TODO(), filter, options.Find().SetProjection(projection))
	if err != nil {
		fmt.Println("Query error:", err)
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, context.TODO())

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

// GetExchangeRates handles the request to retrieve exchange rates.
// @Summary Get exchange rates
// @Description Retrieves the latest exchange rates from an external API.
// @Tags Exchange
// @Accept  json
// @Produce  json
// @Param base query string false "Base currency (default is USD)"
// @Success 200 {object} ExchangeRatesResponse
// @Failure 500 {object} map[string]string "Failed to fetch or decode data"
// @Router /exchange [get]
func GetExchangeRates(w http.ResponseWriter, r *http.Request) {
	baseCurrency := r.URL.Query().Get("base")
	if baseCurrency == "" {
		baseCurrency = "USD"
	}

	url := fmt.Sprintf("https://v6.exchangerate-api.com/v6/4948ebd69671139c0adbb333/latest/%s", baseCurrency)

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Error making GET request", http.StatusInternalServerError)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("Error: Status code %d", resp.StatusCode), http.StatusInternalServerError)
		return
	}

	var data ExchangeRatesResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		http.Error(w, "Error decoding JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		return
	}
}
