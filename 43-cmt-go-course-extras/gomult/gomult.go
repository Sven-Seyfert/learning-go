package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := range [5]int{} {
		time.Sleep(6 * time.Second)
		fmt.Println(i, s)
	}
}

func main() {
	for range [30]int{} {
		go say("bite me!")
	}

	time.Sleep(32 * time.Second)
}
