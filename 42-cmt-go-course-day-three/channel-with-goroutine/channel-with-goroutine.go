package main

import (
	"fmt"
)

func provide(c chan string) {
	c <- "foo"
	c <- "bar"
	c <- "baz"
}

func main() {
	c := make(chan string)
	go provide(c)

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	// fmt.Println(<-c) // At this point it would be throw a fatal error.
}
