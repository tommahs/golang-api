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
func splitstring(url string) ([]string) {
  data := strings.Split(url, "/")
  return data
}

func ErrorWIP() (string)  {
  return "This area is a WIP"
}

func database(urlpath string) () {
  url := splitstring(urlpath[:])
  fmt.Println("split strings result in: ", url)
  fmt.Println("split strings length:",len(url))
  if url[4] != "" {
    if url[3] == "select" {
      fmt.Println("Router send towards /select")
      fmt.Println(mariadb.Selectevent(string(url[4])))
      fmt.Println("done")
    }
  } else {
    fmt.Println("database errror")
  }
}
func Router (w http.ResponseWriter, r *http.Request) {
  fmt.Println("Client attempt:", r.RemoteAddr, "with url: ", r.URL.Path[:])
  if strings.Contains(r.URL.Path[:], "api/db/") {
    database(r.URL.Path[:])
    } else {
    fmt.Fprint(w, "This is an invalid address")
    fmt.Println(r.RemoteAddr, "accessed ", r.URL.Path[4:])
  }
    // fmt.Println(string(r.URL.Path[8:]))


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
