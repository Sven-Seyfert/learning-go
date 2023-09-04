package main

import (
	"bytes"
	"encoding/json"

	"github.com/davecgh/go-spew/spew"
)

// IntOrString is an int that may be unmarshaled from either a JSON number
// literal, or a JSON string.
type IntOrString int

func (i *IntOrString) UnmarshalJSON(d []byte) error {
	var v int
	err := json.Unmarshal(bytes.Trim(d, `"`), &v)
	*i = IntOrString(v)
	return err
}

func main() {
	data := []byte(`[123,"321"]`)
	x := make([]IntOrString, 0)
	if err := json.Unmarshal(data, &x); err != nil {
		panic(err)
	}
	spew.Dump(x)
}
