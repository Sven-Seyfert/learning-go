package main

import (
	"fmt"
)

const (
	f float64 = iota
	g float64 = iota
	h float64 = iota
)

const (
	_  uint32 = iota
	KB uint32 = 1 << (10 * iota)
	MB uint32 = 1 << (10 * iota)
	GB uint32 = 1 << (10 * iota)
)

func main() {
	fmt.Printf("%5.3f\n", f)
	fmt.Printf("%5.3f\n", g)
	fmt.Printf("%5.3f\n", h)

	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
}
