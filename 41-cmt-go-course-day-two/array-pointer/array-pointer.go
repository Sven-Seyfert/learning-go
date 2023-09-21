package main

import (
	"fmt"
	"time"
)

func objdummies() [3]string {
	var o [3]string = [...]string{"foo", "bar", "baz"}
	go func() {
		time.Sleep(2 * time.Second)
		o[1] = "BAR"
	}()
	return o
}

func ptrdummies() *[3]string {
	var o [3]string = [...]string{"FOO", "BAR", "BAZ"}
	go func() {
		time.Sleep(2 * time.Second)
		o[1] = "Bar"
	}()
	return &o
}

func main() {
	a := objdummies()
	fmt.Printf("%v (%T)\n", a, a)
	fmt.Println(a)

	b := ptrdummies()
	fmt.Printf("%v (%T)\n", b, b)
	fmt.Printf("%v (%T)\n", *b, *b)
	fmt.Println(*b)

	time.Sleep(5 * time.Second)
	fmt.Printf("%v (%T)\n", a, a) // hat sich nicht verÃ¤ndert
	fmt.Printf("%v (%T)\n", b, b) // hat sich verÃ¤ndert
}
