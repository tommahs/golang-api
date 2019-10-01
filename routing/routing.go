package routing
import (
  "fmt"
  "net/http"
  "strings"
  "golang-api/handlers/mariadb"
  "golang-api/handlers/filehandler"
)

type event struct {
  id, eventid int
  eventname string
}
func splitstring(url string) ([]string) {
  data := strings.Split(url, "/")
  return data
}

func ErrorWIP() (string)  {
  return "This area is a WIP"
}
func Router (w http.ResponseWriter, r *http.Request) {
  var result string
  var err error
  fmt.Println("Client attempt:", r.RemoteAddr, "with url: ", r.URL.Path[:])
  if strings.Contains(r.URL.Path[:], "api/db/") {
    result,err = mariadb.Main(w, r.URL.Path[8:])
    if err != nil {
      filehandler.Main("errorlog", err)
    }
    // database(r.URL.Path[:])
    } else {
    fmt.Fprint(w, "This is an invalid address")
    fmt.Println(r.RemoteAddr, "accessed ", r.URL.Path[4:])
  }
  fmt.Fprintf(w, result)
}
