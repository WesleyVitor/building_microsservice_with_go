package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {

}

func NewHello() *Hello {
	return &Hello{}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Running Hello Handle")

	content, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	

	fmt.Fprintf(w, "Hello, %s!", content)
}