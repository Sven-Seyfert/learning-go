package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	fmt.Println()

	for _, v := range pow {
		fmt.Printf("%d\n", v)
	}
	fmt.Println()

	for i := range pow {
		fmt.Printf("2**%d\n", i)
	}
	fmt.Println()

	for range pow {
		fmt.Printf("2\n")
	}
}
