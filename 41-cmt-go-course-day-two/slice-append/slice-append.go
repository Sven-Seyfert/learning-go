package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 11)
	printSlice(s)

	s = append(s, 22, 33, 44) // Kapazität wächst mehr
	printSlice(s)

	s = append(s, 55, 66, 77) // Kapazität wächst mehr
	printSlice(s)

	s = append(s, 88, 99, 111) // Kapazität wächst mehr
	printSlice(s)
}
