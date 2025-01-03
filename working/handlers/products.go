package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/WesleyVitor/data"
	"github.com/gorilla/mux"
)

type Products struct {

}

func NewProducts() *Products {
	return &Products{}
}



func (p *Products) GetProducts(w http.ResponseWriter, _ *http.Request) {
	data_products := data.GetProducts()

	enconder := json.NewEncoder(w)
	err := enconder.Encode(data_products)
	
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	
	product := r.Context().Value(KeyProduct{}).(*data.Product)

	data.AddProduct(product)
}

func (p *Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		return
	}

	product := r.Context().Value(KeyProduct{}).(*data.Product)

	err = data.UpdateProduct(id, product)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

}

type KeyProduct struct {}

func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		product := &data.Product{}

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(product)
		if err != nil {
			log.Println("Error deserializing product", err)
			http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
			return
		}

		err = product.Validate()
		if err != nil {
			log.Println("Error validate product", err)
			http.Error(w, fmt.Sprintf("Error to validating json: %s", err), http.StatusBadRequest)
			return
		}
		
		ctx := context.WithValue(r.Context(), KeyProduct{}, product)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}