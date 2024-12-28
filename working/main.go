package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/WesleyVitor/handlers"
)

func main(){


	hello_handler := handlers.NewHello()
	goodbye_handler := handlers.NewGoodbye()

	sm := http.NewServeMux()
	sm.Handle("/", hello_handler)
	sm.Handle("/goodbye", goodbye_handler)

	server := http.Server{
		Addr:         ":9090",      // configure the bind address
		Handler:      sm,                // set the default handler
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	go func ()  {
		log.Println("Server started on: http://localhost:9090")
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()
	
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	sig := <-sigChan
	log.Println("Received terminate, graceful shutdown", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(ctx)

}