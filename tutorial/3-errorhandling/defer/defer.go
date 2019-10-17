package main

import "fmt"

func deferfunction() {
	defer fmt.Println("Done!")
	fmt.Println("Doing some stuff, who knows what?")
}

func main() {
	deferfunction()
}
