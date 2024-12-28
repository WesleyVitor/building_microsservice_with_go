package handlers

import (
	"encoding/json"
	"net/http"

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