package main

import "fmt"

func main() {
	// Warn: These both behave differently.
	// var pow = [...]int{1, 2, 4, 8, 16}
	var pow = []int{1, 2, 4, 8, 16}

	for i, v := range pow {
		if i <= 4 {
			pow[4] = 97
			// v = 97 // Changes only the copy.
		}
		fmt.Printf("2**%d = %d\n", i, v)
	}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
