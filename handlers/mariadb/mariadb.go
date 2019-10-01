package mariadb

import (
  "net/http"
  "golang-api/handlers/mariadb/mariadbrouter"
)


func Main(w http.ResponseWriter, url string) (result string, err error){
  result, err  = mariadbrouter.Main(w,url)
  return result,err
}
