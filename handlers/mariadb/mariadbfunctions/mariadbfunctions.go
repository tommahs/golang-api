package mariadbfunctions

import (
  "fmt"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  mariadbtypes "golang-api/handlers/mariadb/types"
)

func Readevent(table string, data string) (event mariadbtypes.Event, err error) {
  db, err := sql.Open("mysql", "golanguser:golang@/golang")
  defer db.Close()
  command := fmt.Sprintf("SELECT Eventid, Eventname FROM %s WHERE Eventid = %s", string(table), data)
  fmt.Println(command)
  err = db.QueryRow(command).Scan(event.Eventid, event.Eventname)
  results, err := db.Query(command)

  for results.Next(){
      err = results.Scan(&event.Eventid, &event.Eventname)
      if err != nil {
        fmt.Println("SQL error \n", err)
      }
    }
  return event, err
}
