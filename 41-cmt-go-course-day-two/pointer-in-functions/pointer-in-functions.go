package main

import "fmt"

type Vertex struct {
	X, Y int
}

func dontsetit(v Vertex) {
	v.X, v.Y = 33, 44
}

func setit(p *Vertex) {
	p.X, p.Y = 55, 66
}

func main() {
	var v = Vertex{1, 2}
	fmt.Println(v)

	dontsetit(v)
	fmt.Println(v)

	setit(&v)
	fmt.Println(v)
}
