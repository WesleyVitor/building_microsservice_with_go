package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/WesleyVitor/handlers"
)

func main(){
	l := log.New(os.Stdout, "product-api#", log.LstdFlags)

	
	helloHandler := handlers.NewHello(l)
	goodbyeHandler := handlers.NewGoodBye(l)

	routes := http.NewServeMux()
	routes.Handle("/",helloHandler)
	routes.Handle("/goodbye", goodbyeHandler)

	server := &http.Server{
		Addr: ":9000",
		Handler: routes,
		IdleTimeout: 150 * time.Second,
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func(){
		err := server.ListenAndServe()
		if err != nil{
			l.Println(err)
		}
	}()

	sigChan := make(chan os.Signal,1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, syscall.SIGTERM)
	
	sig :=<- sigChan
	l.Printf("\nReceive terminated, graceful shutdown %s\n",sig)
	ctx,cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer func ()  {
		l.Printf("Terminar db")
		cancel()
	}()
	server.Shutdown(ctx)




}