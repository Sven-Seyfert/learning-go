package main

import (
	"fmt"
	"time"
)

func fibonacci(c chan int, quit chan bool) {
	a, b := 0, 1

	for {
		select {
		case c <- a:
			a, b = b, a+b
		case <-quit:
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan bool)

	go fibonacci(c, quit)

	for i := 0; i < 20; i++ {
		fmt.Printf("%v ", <-c)
	}

	quit <- true
	fmt.Printf("...\n")

	time.Sleep(1 * time.Second) // To complete to goroutine.
}
