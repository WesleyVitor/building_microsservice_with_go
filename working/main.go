package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/WesleyVitor/handlers"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main(){

	product_handler := handlers.NewProducts()

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products", product_handler.GetProducts)
	
	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/products/{id:[0-9]+}", product_handler.UpdateProduct)	
	putRouter.Use(product_handler.MiddlewareProductValidation)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/products", product_handler.AddProduct)
	postRouter.Use(product_handler.MiddlewareProductValidation)
	
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	server := http.Server{
		Addr:         ":9090",      // configure the bind address
		Handler:      ch(sm),       // set the default handler
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	go func ()  {
		log.Println("Server started on: http://localhost:9090")
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("Error starting server: %s\n", err)
			os.Exit(1) // Return os.Interrupt
		}
	}()
	
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	sig := <-sigChan
	log.Println("Received terminate, graceful shutdown", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(ctx)

}