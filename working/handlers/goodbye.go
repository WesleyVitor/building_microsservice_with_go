package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type Goodbye struct{
	l *log.Logger
}

func NewGoodBye(l *log.Logger) *Goodbye{
	return &Goodbye{l}
}


func (goodbye *Goodbye) ServeHTTP(rw http.ResponseWriter, req *http.Request){
	goodbye.l.Printf("Acessou %s\n", req.URL)
	fmt.Fprintf(rw,"Goodbye\n")
}