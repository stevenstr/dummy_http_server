package main

import (
	"log"
	"net/http"
)

func mainPage(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Welcome buddy!"))
}

func apiPage(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("welcome to the api page!"))
}

func main() {
	log.Println("dummy service is up!")

	http.HandleFunc("/", mainPage)
	http.HandleFunc("/api", apiPage)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
