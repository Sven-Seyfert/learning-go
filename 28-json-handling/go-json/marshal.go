package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	x := map[string]string{
		"foo": "bar",
	}
	data, _ := json.Marshal(x)
	fmt.Println(string(data))
}
