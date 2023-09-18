package main

import "fmt"

func showit(s string) {
	fmt.Printf("[%v]\n", s)
}

func main() {
	showit("Hello from Go!")
	showit("Er fragte: \"Wie geht's?\"")
	showit("erstens\nzweitens")
	showit("erstens\tzweitens")
	showit(`\"\n\`) // "raw" strings
	showit("Es kostet \u20ac 100,-")
	showit("Es kostet \U000020ac 100,-")
	showit("Es kostet \xe2\x82\xac 100,-")
	showit("Es kostet \xe2\x82 100,-") // ungÃ¼ltiges UTF-8
}
