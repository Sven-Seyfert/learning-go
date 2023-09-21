package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter text: ")
	var t string
	fmt.Scanf("%s", &t)
	fmt.Println(t)

	fmt.Print("Enter number: ")
	var i int
	fmt.Scanf("%d", &i) // Warn: The rest is still in the quue.
	fmt.Println(i)
}
