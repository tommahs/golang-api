package mariadbfunctions

import (
	"database/sql"
	"fmt"
	mariadbtypes "golang-api/example-api/handlers/mariadb/types"

	_ "github.com/go-sql-driver/mysql"
)

func Readevent(table string, data string) (event mariadbtypes.Event, err error) {
	db, err := sql.Open("mysql", "golanguser:golang@/golang")
	defer db.Close()
	command := fmt.Sprintf("SELECT Eventid, Eventname FROM %s WHERE Eventid = %s", string(table), data)
	results, err := db.Query(command)

	for results.Next() {
		err = results.Scan(&event.Eventid, &event.Eventname)
		if err != nil {
			fmt.Println("SQL error \n", err)
		}
	}
	return event, err
}
