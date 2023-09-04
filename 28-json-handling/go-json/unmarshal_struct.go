package main

import (
	"encoding/json"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	type person struct {
		Name        string `json:"name"`
		Age         int    `json:"age"`
		Description string `json:"descr,omitempty"`
		secret      string // Unexported fields are never (un)marshaled
	}
	data := []byte(`{"name":"Bob","age":32,"secret":"Shhh!"}`)
	var x person
	_ = json.Unmarshal(data, &x)
	spew.Dump(x)
}
