package main

import (
	"fmt"
	"strings"
)

var data = `---
conn:
    database: sandbox
    host:     "%s"
    port:     5432
...
`

var attack = `foreignserver"
    port:     5431
otherconn:
    host:     "someserver`

func main() {
	fmt.Println("= Unsecure variant")
	fmt.Printf(data, attack)

	data = strings.ReplaceAll(data, "\"%s\"", "%q")

	fmt.Println("= Attack catched")
	fmt.Printf(data, attack)
}
