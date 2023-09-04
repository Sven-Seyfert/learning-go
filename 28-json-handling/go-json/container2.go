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

type Responses []interface{}

func (r *Responses) UnmarshalJSON(d []byte) error {
	var raw []json.RawMessage
	if err := json.Unmarshal(d, &raw); err != nil {
		return err
	}
	var success Success
	var fail Error
	result := make(Responses, len(raw))
	for i, msg := range raw {
		_ = json.Unmarshal(msg, &success)
		if success.Status == "ok" {
			result[i] = success
			continue
		}
		_ = json.Unmarshal(msg, &fail)
		result[i] = fail
	}
	*r = result
	return nil
}
// OMIT

func main() {
	data := []byte(`[
		{"status":"ok","results": ["one", "two"]},
		{"status":"not found", "reason": "The requested object does not exist"}
	]`)
	var x Responses
	if err := json.Unmarshal(data, &x); err != nil {
		panic(err)
	}
	spew.Dump(x)
}
