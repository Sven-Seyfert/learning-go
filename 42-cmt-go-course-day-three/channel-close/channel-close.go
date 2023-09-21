package main

import "fmt"

func collatz(n int, c chan int) {
	defer func() {
		close(c)
	}()

	for {
		c <- n
		if n == 1 {
			return
		}

		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
}

func main() {
	c := make(chan int)
	go collatz(11, c)

	// Variant 1: "ok" has to be checked manuelly
	// for {
	// 	v, ok := <-c
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println(v)
	// }

	// Variant 2: Range brings the "ok" out of the box.
	// Recommended way of doing it.
	for v := range c {
		fmt.Println(v)
	}
}
