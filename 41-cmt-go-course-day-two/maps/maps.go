package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])
	fmt.Println("The map:", m)

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	m["Answer"] = 42
	v, ok = m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	m["Answer"] = 0
	v, ok = m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
