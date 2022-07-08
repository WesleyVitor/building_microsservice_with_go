package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct{
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello{
	return &Hello{l}
}

func (hello *Hello) ServeHTTP(rw http.ResponseWriter, req *http.Request){
	hello.l.Printf("Acessou %s\n", req.URL)
	content, err := ioutil.ReadAll(req.Body)
	if err!=nil{
		http.Error(rw,"Opps",http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw,"Hello %s\n", string(content))
}
