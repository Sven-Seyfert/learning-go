package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		fmt.Println("bye-bye!")
	}()

	os.Exit(1)
}
