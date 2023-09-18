package main

import "fmt"

func main() {
aussen:
	for {
		for {
			break aussen
		}
	}
	fmt.Println("fertig.")
}
