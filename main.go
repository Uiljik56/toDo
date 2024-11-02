package main

import (
	"log"
	"net/http"
)

func main() {

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		log.Println("привет")
	}

	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
