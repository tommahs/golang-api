package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var testwaitgroup sync.WaitGroup

func Testcleanup() {
	if r := recover(); r != nil {
		fmt.Println("Cleanup wordt uitgevoerd omdat", r)
	}
	defer testwaitgroup.Done()
}

func Testsay(s string) {
	defer Testcleanup()
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		if i == 2 {
			panic(s + " een twee tegenkomt")
		}
		fmt.Println(s, "- optellen:", i)
	}
}

func TestMain(t *testing.T) {
	// var testwaitgroup sync.WaitGroup
	testwaitgroup.Add(1)
	go Testsay("coroutine 1")
	testwaitgroup.Add(1)
	go Testsay("coroutine 2")
	testwaitgroup.Wait()
}
