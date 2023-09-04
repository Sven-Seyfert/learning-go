package main

import (
	"encoding/json"

	"github.com/davecgh/go-spew/spew"
)

type Success struct {
	Results []string `json:"results"`
}

type Error struct {
	Error  string `json:"error"`
	Reason string `json:"reason"`
}

type Response struct {
	Success
	Error
}

func main() {
	data := []byte(`[
		{"results": ["one", "two"]},
		{"error":"not found", "reason": "The requested object does not exist"}
	]`)
	x := make([]Response, 0)
	if err := json.Unmarshal(data, &x); err != nil {
		panic(err)
	}
	spew.Dump(x)
}
