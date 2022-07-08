package main

import (
	"log"
	"net/http"
	"os"

	"github.com/WesleyVitor/handlers"
)

func main(){
	l := log.New(os.Stdout, "product-api#", log.LstdFlags)

	
	helloHandler := handlers.NewHello(l)
	goodbyeHandler := handlers.NewGoodBye(l)

	server := http.NewServeMux()
	server.Handle("/",helloHandler)
	server.Handle("/goodbye", goodbyeHandler)

	
	http.ListenAndServe(":9000", server)

}