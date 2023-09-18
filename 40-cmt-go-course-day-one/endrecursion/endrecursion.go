package main

import "fmt"

func factorial_zweit(total, n int) int {
	if n == 0 {
		return total
	} else {
		return factorial_zweit(total*n, n-1)
	}
}

func factorial_endrek(n int) int {
	return factorial_zweit(1, n)
}

func main() {
	fmt.Println(factorial_endrek(6))
}
