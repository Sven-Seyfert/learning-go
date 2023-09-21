package main

import "fmt"

func main() {
	exampleOne()
	fmt.Println()

	exampleTwo()
	fmt.Println()

	exampleThree()
	fmt.Println()
}

func exampleOne() {
	i, j := 42, 2701

	p := &i
	fmt.Printf("%v %p %T\n", p, p, p)
	fmt.Println(*p)

	*p = 21
	fmt.Println(i)

	p = &j
	fmt.Printf("%v %p %T\n", p, p, p)

	*p /= 37
	fmt.Println(j)
}

func exampleTwo() {
	var i int
	var p *int
	fmt.Printf("%v (%T)\n", p, p)
	p = &i
	fmt.Printf("%v (%T)\n", p, p)
}

func exampleThree() {
	// Pointer with new
	var p *int
	p = new(int)

	fmt.Printf("p: %v an %p (%T)\n", *p, p, p)
}
