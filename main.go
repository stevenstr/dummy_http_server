package main

import (
	"log"
	"net/http"
)

func mainHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Welcome buddy!"))
}

func apiHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("welcome to the api page!"))
}

func apiAuthHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Welcome to api/auth page..."))
}

func clientHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Welcome to client page!"))
}

func main() {
	log.Println("dummy service is up!")

	mux := http.NewServeMux()

	mux.HandleFunc("/", mainHandler)
	mux.HandleFunc("/client/", clientHandler)
	mux.HandleFunc("/api/", apiHandler)
	mux.HandleFunc("/api/auth", apiAuthHandler)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
