package main

import (
	"log"
	"me/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/about", handlers.HandleAboutInfo)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
