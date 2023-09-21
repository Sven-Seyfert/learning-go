package main

import (
	"fmt"
	"regexp"
)

var input = `---
conn:
    database: sandbox
    user:     @@jdoe@@
    host:     @@host@@
    port:     @@port@@
...`

func main() {
	host := "myserver"
	port := 5432

	r := regexp.MustCompile(`(@@)(\w+)(@)(@)`)
	result := r.ReplaceAllStringFunc(input, func(m string) string {
		parts := r.FindStringSubmatch(m)
		fmt.Printf("m:     %v (%T)\n", m, m)
		fmt.Printf("parts: %v (%T)\n", parts, parts)

		switch parts[2] {
		case "host":
			return fmt.Sprintf("%q", host)
		case "port":
			return fmt.Sprintf("%d", port)
		default:
			return parts[0]
		}
	})

	fmt.Println(result)
}
