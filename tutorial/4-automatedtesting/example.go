package main

import (
	"fmt"
	"sync"
	"time"
)

var waitgroup sync.WaitGroup

func cleanup() {
	if r := recover(); r != nil {
		fmt.Println("Cleanup wordt uitgevoerd omdat", r)
	}
	defer waitgroup.Done()
}

func say(s string) {
	defer cleanup()
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		if i == 2 {
			panic(s + " een twee tegenkomt")
		}
		fmt.Println(s, "- optellen:", i)
	}
}

func main() {
	waitgroup.Add(1)
	go say("coroutine 1")
	waitgroup.Add(1)
	go say("coroutine 2")
	waitgroup.Wait()
}
