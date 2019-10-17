package mariadbrouter

import (
	"golang-api/handlers/mariadb/mariadbfunctions"
	mariadbtypes "golang-api/handlers/mariadb/types"
	"net/http"
	"strings"
)

func Main(w http.ResponseWriter, url string) (result string, err error) {
	var table string
	var method string
	var data string
	var event mariadbtypes.Event
	splittedurl := strings.Split(url, "/")
	if len(splittedurl) < 3 {
		return "Invalid call\n", err
	}
	table = splittedurl[0]
	method = splittedurl[1]
	data = splittedurl[2]
	if table == splittedurl[0] {
		switch method {
		case "read":
			if data == "" {
				result = "Error in value\n"
				return result, err
			} else {
				event, err = mariadbfunctions.Readevent(table, data)
				if err != nil {
					result = "error in reading data\n"
					return result, err
				} else {
					result = "Selected data from event: \nID = " + string(event.Eventid) + "\nName = " + string(event.Eventname)
					return result, err
				}
			}
		default:
			result = "Default reaction"
			return result, err
		}
	} else {
		result = "No table"
		return result, err
	}
	result = "No route"
	return result, err
}
