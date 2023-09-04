package main

import (
	"encoding/json"

	"github.com/davecgh/go-spew/spew"
)

// SliceOrString is a slice of strings that may be unmarshaled from either
// a JSON array of strings, or a single JSON string.
type SliceOrString []string

func (s *SliceOrString) UnmarshalJSON(d []byte) error {
	if d[0] == '"' {
		var v string
		err := json.Unmarshal(d, &v)
		*s = SliceOrString{v}
		return err
	}
	var v []string
	err := json.Unmarshal(d, &v)
	*s = SliceOrString(v)
	return err
}

func main() {
	data := []byte(`["one", ["two","elements"]]`)
	x := make([]SliceOrString, 0)
	if err := json.Unmarshal(data, &x); err != nil {
		panic(err)
	}
	spew.Dump(x)
}
