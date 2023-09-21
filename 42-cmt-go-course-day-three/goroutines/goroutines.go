package main

import (
	"fmt"
	"time"
)

func say(s string, n uint) {
	for i := uint(0); i < n; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i, s)
	}
}

func main() {
	go func() {
		say("world", 10) // Not all ten times will be reached, because the routine can not be ended. WaitGroup is necessary.
	}()

	say("hello", 5)
	// time.Sleep(2 * time.Second)
}
