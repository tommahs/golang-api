package main

import (
	"fmt"
	"golang-api/tutorial/1-api/router"
)

func main() {
	var result = router.Router("Succesfull import")
	fmt.Println(result)
}
