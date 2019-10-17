package routing

import (
	"fmt"
	"golang-api/handlers/filehandler"
	"golang-api/handlers/mariadb"
	"net/http"
	"strings"
)

func Router(w http.ResponseWriter, r *http.Request) {
	var err error
	var result string
	fmt.Println("Client attempt:", r.RemoteAddr, "with url: ", r.URL.Path[:])
	if strings.Contains(r.URL.Path[:], "api/db/") {
		result, err = mariadb.Main(w, r.URL.Path[8:])
		if err != nil {
			filehandler.Main("errorlog", err)
		}
		fmt.Fprint(w, result)
		// database(r.URL.Path[:])
	} else {
		fmt.Fprint(w, "Succesfull connection with the API!\n")
		fmt.Println(r.RemoteAddr, "accessed ", r.URL.Path[4:])

	}
	fmt.Fprintf(w, "\nApi works but no logic for this URL...\n")
}
