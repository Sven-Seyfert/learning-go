package main

import "fmt"

func main() {
	stringMid()
	fmt.Println()
	stringMidByRune()
}

func stringMid() {
	s := "Hallöchen"

	fmt.Println(s)
	fmt.Println(s[4:6])
	fmt.Println(s[4:5])
	fmt.Println(s[:4])
	fmt.Println(s[5:])
}

func stringMidByRune() {
	s := "Hallöchen"
	r := []rune(s)

	println(string(r))
	println(string(r[4:6]))
	println(string(r[4:5]))
	println(string(r[:4]))
	println(string(r[5:]))
}
