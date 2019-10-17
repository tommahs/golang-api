package main

import (
	"fmt"
	"sync"
	"time"
)

var waitgroup sync.WaitGroup

func say(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
	waitgroup.Done()
}

func example4() {
	waitgroup.Add(1)
	go say("Functie 1 - coroutine")
	waitgroup.Add(1)
	go say("Functie 2 - coroutine")
	waitgroup.Wait()
}

func main() {
	example4()
}
