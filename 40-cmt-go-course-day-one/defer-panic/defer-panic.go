package main

import (
	"fmt"
	"log"
)

func forcePanic() {
	var a []int = make([]int, 0)
	fmt.Println(a[2])
}

func main() {
	// Panic will not skip defer
	defer fmt.Println("Have to clean up...")
	forcePanic()
	fmt.Println("Bye!")

	// Fatal will skip defer
	log.Fatal("")
}
