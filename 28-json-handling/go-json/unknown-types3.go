package main

import (
	"encoding/json"

	"github.com/davecgh/go-spew/spew"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Animal struct {
	Species string `json:"species"`
	Name    string `json:"name"`
}

type Address struct {
	City   string `json:"city"`
	Street string `json:"street"`
}

type Responses []interface{}
// OMIT

func (r *Responses) UnmarshalJSON(d []byte) error {
	var tmp []json.RawMessage
	if err := json.Unmarshal(d, &tmp); err != nil {
		return err
	}
	var header struct {
		Type string `json:"type"`
	}
	result := make(Responses, len(tmp))
	// MID OMIT
	for i, raw := range tmp {
		_ = json.Unmarshal(raw, &header)
		switch header.Type {
		case "person":
			tgt := Person{}
			_ = json.Unmarshal(raw, &tgt)
			result[i] = tgt
		case "animal":
			tgt := Animal{}
			_ = json.Unmarshal(raw, &tgt)
			result[i] = tgt
		case "address":
			tgt := Address{}
			_ = json.Unmarshal(raw, &tgt)
			result[i] = tgt
		}
	}
	*r = result
	return nil
}
// OMIT

func main() {
	data := []byte(`[
		{"type": "person", "name": "Bob", "age": 32},
		{"type": "animal", "species": "dog", "name": "Spot"},
		{"type": "address", "city": "Rotterdam", "street": "Goudsesingel"}
	]`)
	var x Responses
	if err := json.Unmarshal(data, &x); err != nil {
		panic(err)
	}
	spew.Dump(x)
}
