package main

import "fmt"

type vertex struct {
	x int
	y int
}

func main() {
	v := vertex{1, 2}
	p := &v

	(*p).x = 20
	p.x = 20 // gleichwertige Schreibweise

	fmt.Println(v)
}
