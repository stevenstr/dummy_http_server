package main

import (
	"log"
	"net/http"
)

type Handler struct{}

func (h Handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	data := []byte("Hello world!")
	res.Write(data)
}

func main() {
	log.Println("dummy service is up!")
	var h Handler
	if err := http.ListenAndServe(":8080", h); err != nil {
		log.Fatal(err)
	}
}
