package main

import (
	"fmt"
	"time"
)

func sleep_and_signal(n int, c chan bool) {
	time.Sleep(time.Duration(n) * time.Second)
	c <- true
}

func main() {
	fmt.Println("Start")

	c := make(chan bool)
	go sleep_and_signal(3, c)

loop:
	for {
		select {
		case <-c:
			fmt.Println("A message!")
			break loop
		default:
			fmt.Println("Nothing!")
			time.Sleep(500 * time.Millisecond)
		}
	}

	fmt.Println("Done")
}
