package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	example()
	fmt.Println()

	structNewExample()
}

func example() {
	fmt.Println(Vertex{}) // valid
	fmt.Println(Vertex{1, 2})
	fmt.Println(Vertex{0, 1})
	// fmt.Println(Vertex{1}) // invalid
	fmt.Println(Vertex{X: 1})
	fmt.Println(Vertex{Y: 55})
	fmt.Println(Vertex{
		X: 0,
		Y: 1,
	})
}

func structNewExample() {
	var pv *Vertex

	pv = new(Vertex)
	fmt.Println(pv)

	pv = &Vertex{} // impliziert new()
	fmt.Println(pv)

	pv = &Vertex{1, 2}
	fmt.Println(pv)
}
