package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Printf("%v (%T)\n", primes, primes)

	var s []int = primes[1:4]
	fmt.Printf("%v (%T)\n", s, s)
	fmt.Println(s)

	fmt.Println()
	boundaries()
}

func boundaries() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}
