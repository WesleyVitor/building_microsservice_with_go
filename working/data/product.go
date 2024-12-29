package data

import (
	"errors"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

var ErrProductNotFound = errors.New("Product not found")

func UpdateProduct(id int, p *Product) error{
	_, pos := findProduct(id)
	if pos == -1 {
		return ErrProductNotFound
	}
	p.ID = id
	productList[pos] = p
	return nil
}

func findProduct(id int) (*Product, int) {
	for i, p := range productList {
		if p.ID == id {
			return p, i
		}
	}
	return nil, -1
}

func GetProducts() []*Product {
	return productList
}

func AddProduct(p *Product) {

	p.ID = getNextID()
	productList = append(productList, p)
}

func getNextID() int {
	if len(productList) == 0 {
		return 1
	}
	lp := productList[len(productList)-1]
	
	return lp.ID + 1
}

var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}