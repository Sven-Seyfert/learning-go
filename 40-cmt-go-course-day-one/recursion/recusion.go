package main

import "fmt"

func factorial_recursive(n int) int {
	if n <= 1 {
		return n
	}
	return n * factorial_recursive(n-1)
}

func factorial_loop(n int) int {
	r := 1
	for i := 1; i <= n; i += 1 {
		r *= i
	}
	return r
}

func main() {
	fmt.Println(factorial_recursive(6000))
	fmt.Println(factorial_loop(6))
}
