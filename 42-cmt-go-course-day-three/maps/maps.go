package main

import "fmt"

func main() {
	a := wordUnique([]string{"foo", "bar", "foo", "baz", "foo", "bar"})
	fmt.Println(a)

	b := wordCount([]string{"foo", "bar", "foo", "baz", "foo", "bar"})
	fmt.Println(b)
}

func wordUnique(s []string) map[string]bool {
	m := make(map[string]bool)
	for _, f := range s {
		m[f] = true
	}

	return m
}

func wordCount(s []string) map[string]int {
	m := make(map[string]int)
	for _, f := range s {
		m[f]++
	}

	return m
}
