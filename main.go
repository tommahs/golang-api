package main

import (
	"fmt"
	"golang-api/greet"
	"golang-api/routing"
	"log"
	"net/http"
)

// variable definitions
var hostname = "localhost:8080"

func main() {
	fmt.Println(greet.Morning)
	http.HandleFunc("/", routing.Router)
	log.Fatal(http.ListenAndServe(hostname, nil))
}
