package main

import (
  "fmt"
  "log"
  "net/http"
  "golang-api/greet"
  "golang-api/routing"
)


// variable definitions
var hostname = "localhost:8080"

func main() {
  fmt.Println(greet.Morning)
  http.HandleFunc("/", routing.Router)
  log.Fatal(http.ListenAndServe(hostname, nil))
}
