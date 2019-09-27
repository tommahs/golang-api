package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "golang-api/greet"
  "golang-api/routing"
  // "golang-api/src/routing"
)
// type definitions
type Page struct {
  Title string
  Body []byte
}

// variable definitions
var hostname = "localhost:8080"

// save page
func (p *Page) save () error {
  filename := p.Title + ".text"
  return ioutil.WriteFile(filename, p.Body, 0600)
}

// load page
func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
      return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}


func main() {
  fmt.Println(greet.Morning)
  http.HandleFunc("/", routing.Router)
  log.Fatal(http.ListenAndServe(hostname, nil))
}

// Import all modules into the system (routes/handlers)

func loadModules() {
  

}
