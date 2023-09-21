package main

import (
	"fmt"
	"time"
)

func after(n int) chan bool {
	c := make(chan bool)
	go func(c chan bool) {
		time.Sleep(time.Duration(n) * time.Second)
		c <- true
	}(c)
	return c
}

func timer(n int) chan bool {
	c := make(chan bool)
	go func(c chan bool) {
		for {
			time.Sleep(time.Duration(n) * time.Second)
			c <- true
		}
	}(c)
	return c
}

func main() {
	t := timer(1)
	b := after(5)

	for {
		select {
		case <-t:
			fmt.Print("tick")
		case <-b:
			fmt.Print("BOOM!")
			fmt.Println()
			return
		default:
			fmt.Print(".")
			time.Sleep(time.Duration(50) * time.Millisecond)
		}
	}
}
