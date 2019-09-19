package routing
import (
  "fmt"
  "net/http"
  "golang-api/handlers/mariadb"
  "strings"
)

type event struct {
  id, eventid int
  eventname string
}

//
func Router (w http.ResponseWriter, r *http.Request) {
  fmt.Println("split strings result in: ", strings.Split(r.URL.Path, "/"))
  fmt.Println("Router Access URL-Test: ", r.URL.Path[4:])
  fmt.Println("Client:", r.RemoteAddr, "accessed ", r.URL.Path[4:])
  if strings.Contains(r.URL.Path, "/db") { // WIP
    switch r.URL.Path[7:] {
    case "":
      fmt.Println("database root")
    case "/create":
      fmt.Println(true)
    case "/update":
      fmt.Println(mariadb.Updateevent())
    case "/select":
      fmt.Println(mariadb.Selectevent())
    }
  } else {
    fmt.Fprint(w, "This is an invalid address")
    fmt.Println(r.RemoteAddr, "accessed ", r.URL.Path[4:])
  }


  // switch r.URL.Path[4:] {
  // case "/db":
  //   fmt.Fprint(w, "This sub is for database")
  // case "/testdata":
  //   fmt.Println(r.RemoteAddr, "accessed ", r.URL.Path[4:])
  //   fmt.Fprint(w, "This sub is for testdata")
  // default:
  //
  // }
}
