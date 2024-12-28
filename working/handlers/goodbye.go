package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Goodbye struct {

}

func NewGoodbye() *Goodbye {

	return &Goodbye{}
}

func (g *Goodbye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Running Goodbye Handle")

	content, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Goodbye %s!", content)

}