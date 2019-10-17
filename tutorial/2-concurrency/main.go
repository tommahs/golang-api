package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
func example1() {
	go say("Functie 'say' 1 - coroutine")
	say("Functie 'say' 2")
}
func example2() {

	say("Functie 'say' 1")
	go say("Functie 'say' 2  -coroutine")
}

func example3() {
	go say("Functie 'say' 1 - coroutine")
	go say("Functie 'say' 2  -coroutine")
}
func main() {
	fmt.Println("Example 1")
	example1()

	fmt.Println("\nExample 2")
	example2()

	fmt.Println("\nExample 3")
	example3()
	fmt.Println("\n end 3")
}
