package main

import "fmt"

type Vertex struct {
	X, Y int
}

type Mark struct {
	V     *Vertex // Notice the pointer
	Title string
}

func dontsetit(m Mark) {
	m.V.X, m.V.Y = 33, 44 // Will be set anyway
	m.Title = "bar"
}

func setit(m *Mark) {
	m.V.X, m.V.Y = 55, 66
	m.Title = "baz"
}

func main() {
	var m = Mark{
		V:     &Vertex{1, 2},
		Title: "foo",
	}
	fmt.Println(m, m.V)

	dontsetit(m)
	fmt.Println(m, m.V)

	setit(&m)
	fmt.Println(m, m.V)

	n := m
	fmt.Println(n, n.V)

	m.V.X = 99
	fmt.Println(m, m.V)
	fmt.Println(n, n.V)
}
