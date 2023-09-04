package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type person struct {
		Name        string `json:"name"`
		Age         int    `json:"age"`
		Description string `json:"descr,omitempty"`
		secret      string // Unexported fields are never (un)marshaled
	}
	x := person{
		Name:   "Bob",
		Age:    32,
		secret: "Shhh!",
	}
	data, _ := json.Marshal(x)
	fmt.Println(string(data))
}
