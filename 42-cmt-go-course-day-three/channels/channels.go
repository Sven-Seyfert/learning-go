package main

import (
	"fmt"
)

func main() {
	c := make(chan string, 3)

	c <- "foo"
	c <- "bar"
	c <- "baz"

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
