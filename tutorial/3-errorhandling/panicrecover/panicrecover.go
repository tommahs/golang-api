package main

import (
	"fmt"
	"time"
)

func cleanup() {
	if r := recover(); r != nil {
		fmt.Println("Cleanup wordt uitgevoerd omdat:", r)
	}
}

func say(s string) {
	defer cleanup()
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		if i == 2 {
			panic("We hebben een twee!")
		}
		fmt.Println(s, i)
	}
}

func main() {
	say("We tellen op:")
}
