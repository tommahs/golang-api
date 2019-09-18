package routing
import (
  "fmt"
  "net/http"
  "golang-api/handlers/mariadb"
)

//
func router (w http.ResponseWriter, r *http.Request) {
  fmt.Println("Router Access URL-Test: ", r.URL.Path[4:])
  switch r.URL.Path[4:] {
  case "/db":
    fmt.Println(r.RemoteAddr, "accessed ", r.URL.Path[4:])
    fmt.Fprint(w, "This sub is for database")
  case "/testdata":
    fmt.Println(r.RemoteAddr, "accessed ", r.URL.Path[4:])
    fmt.Fprint(w, "This sub is for testdata")
  default:
    fmt.Fprint(w, "This area is WIP")
    fmt.Println(r.RemoteAddr, "accessed ", r.URL.Path[4:])
  }
}
