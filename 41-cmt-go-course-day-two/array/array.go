package main

import "fmt"

func main() {

	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	var b [3]string = [3]string{"foo", "bar", "baz"}
	fmt.Println(b[0], b[1])
	fmt.Println(b)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	fmt.Printf("%v (%T)\n", primes, primes)
}
