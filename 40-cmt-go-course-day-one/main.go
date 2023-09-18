package main

import (
	"fmt"
	"strings"
)

func main() {
	returnOne, returnTwo := swap("world", "hello")
	fmt.Println(returnOne, returnTwo)
	seperator()

	fmt.Println(split(17))
	seperator()

	// Variadic paremeters
	sum(1, 2)
	sum(1, 2, 3)
	sum()

	nums := []int{1, 2, 3, 4}
	sum(nums...)
	seperator()

	swap2()
	seperator()

	fmt.Println(reverse("hello"))
	seperator()
}

func seperator() {
	fmt.Println(strings.Repeat("-", 40))
}

// Swap returns multiple return values
func swap(first, second string) (string, string) {
	return second, first
}

// Named return values (not recommended)
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// Sum uses a variadic parameter
func sum(nums ...int) {
	fmt.Print("Sum of ", nums, " = ")

	total := 0
	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}

// Swap2 shows how variable values can be swaped directly.
func swap2() {
	a, b := 33, 444
	fmt.Println(a, b)

	b, a = a, b
	fmt.Println(a, b)
}

// Reverse is made by Bertram Scharpf
func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
