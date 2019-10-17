package mariadb

import (
	"golang-api/example-api/handlers/mariadb/mariadbrouter"
	"net/http"
)

func Main(w http.ResponseWriter, url string) (result string, err error) {
	result, err = mariadbrouter.Main(w, url)
	return result, err
}
