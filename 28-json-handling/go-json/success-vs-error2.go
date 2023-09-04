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
	Success
	Error
}

func (r *Response) UnmarshalJSON(d []byte) error {
	if err := json.Unmarshal(d, &r.Success); err != nil {
		return err
	}
	if r.Success.Status == "ok" {
		return nil
	}
	r.Success = Success{}
	return json.Unmarshal(d, &r.Error)
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
