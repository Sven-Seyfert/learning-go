package main

import (
	"encoding/json"

	"github.com/davecgh/go-spew/spew"
)

type Success struct {
	Status  string   `json:"status"`
	Results []string `json:"results"`
}

type Error struct {
	Status string `json:"status"`
	Reason string `json:"reason"`
}

type Response struct {
	Result interface{}
}

func (r *Response) UnmarshalJSON(d []byte) error {
	var success Success
	if err := json.Unmarshal(d, &success); err != nil {
		return err
	}
	if success.Status == "ok" {
		r.Result = success
		return nil
	}
	var fail Error
	if err := json.Unmarshal(d, &fail); err != nil {
		return err
	}
	r.Result = fail
	return nil
}
// OMIT

func main() {
	data := []byte(`[
		{"status":"ok","results": ["one", "two"]},
		{"status":"not found", "reason": "The requested object does not exist"}
	]`)
	x := make([]Response, 0)
	if err := json.Unmarshal(data, &x); err != nil {
		panic(err)
	}
	spew.Dump(x)
}
