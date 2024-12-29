package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/WesleyVitor/data"
)

type Products struct {

}

func NewProducts() *Products {
	return &Products{}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(w, r)
		return
	}

	if r.Method == http.MethodPut {
		log.Println("Update a product")
		p.updateProduct(w, r)
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(w http.ResponseWriter, _ *http.Request) {
	data_products := data.GetProducts()

	enconder := json.NewEncoder(w)
	err := enconder.Encode(data_products)
	
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(w http.ResponseWriter, r *http.Request) {
	product := &data.Product{}
	
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(product)
	if err != nil {
		log.Println("Error deserializing product", err)
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}

	data.AddProduct(product)
}

func (p *Products) updateProduct(w http.ResponseWriter, r *http.Request) {
	log.Println("Update a product", r.URL.Query())
	id := r.URL.Query().Get("id")

	product := &data.Product{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(product)
	if err != nil {
		log.Println("Error deserializing product", err)
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}

	id_int, _ := strconv.Atoi(id)
	err = data.UpdateProduct(id_int, product)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

}