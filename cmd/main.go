package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("dummy service is up!")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
