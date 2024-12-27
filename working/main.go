package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main(){


	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request){
		log.Println("Running Goodbye Handle")

		content, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "Goodbye %s!", content)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		log.Println("Running Hello Handle")

		content, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}


		fmt.Fprintf(w, "Hello, %s!", content)
	
	})

	log.Println("Server started on: http://localhost:9090")

	http.ListenAndServe(":9090", nil)

}