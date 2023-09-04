package main

import (
	"bytes"
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
	dec := json.NewDecoder(bytes.NewReader(d))
	_, _ = dec.Token() // Consume opening "["
	result := make(Responses, 0)
	var header struct {
		Type string `json:"type"`
	}
	// MID OMIT
	var raw json.RawMessage
	for dec.More() {
		_ = dec.Decode(&raw)
		_ = json.Unmarshal(raw, &header)
		switch header.Type {
		case "person":
			tgt := Person{}
			_ = json.Unmarshal(raw, &tgt)
			result = append(result, tgt)
		case "animal":
			tgt := Animal{}
			_ = json.Unmarshal(raw, &tgt)
			result = append(result, tgt)
		case "address":
			tgt := Address{}
			_ = json.Unmarshal(raw, &tgt)
			result = append(result, tgt)
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
