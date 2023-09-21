package main

import "fmt"

func holdSum() func(int) {
	sum := 0

	return func(x int) {
		sum += x
		fmt.Println("Die Summe ist jetzt:", sum)
	}
}

func main() {
	a := holdSum()
	b := holdSum()

	a(3)
	a(4)
	b(100)
	a(7)
	b(200)
}
