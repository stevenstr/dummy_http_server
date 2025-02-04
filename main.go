package main

import (
	"fmt"
	"log"
	"net/http"
)

func mainHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Welcome buddy!\n"))
	body := fmt.Sprintln("Request Method:", req.Method)
	body += fmt.Sprintln("Request headers:")
	body += fmt.Sprintln()
	for k, v := range req.Header {
		body += fmt.Sprintf("%s: \n", k)
		for _, v := range v {
			body += fmt.Sprintf("		%s\n", v)
		}
	}

	if err := req.ParseForm(); err != nil {
		res.Write([]byte(err.Error()))
		return
	}

	body += fmt.Sprintln()
	body += fmt.Sprintln("Request querry:")
	for k, v := range req.Form {
		body += fmt.Sprintf("   %s: %s\n", k, v)
	}

	res.Write([]byte(body))
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
	mux.HandleFunc("GET /client/", clientHandler)
	mux.HandleFunc("GET /api/", apiHandler)
	mux.HandleFunc("GET /api/auth", apiAuthHandler)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
