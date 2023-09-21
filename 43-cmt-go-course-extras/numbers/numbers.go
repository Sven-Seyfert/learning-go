package main

// go build -gcflags='-S -N' numbers.go 2>&1 | vi - +/ADD

import (
	"fmt"
	"math/rand"
)

func main() {
	var i uint32 = rand.Uint32()
	fmt.Println(i)

	i += rand.Uint32()
	fmt.Println(i)

	var x float64 = rand.Float64()
	fmt.Println(x)

	x += rand.Float64()
	fmt.Println(x)
}
