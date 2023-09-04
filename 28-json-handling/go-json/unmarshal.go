package main

import (
	"encoding/json"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	data := []byte(`{"foo":"bar"}`)
	var x interface{}
	_ = json.Unmarshal(data, &x)
	spew.Dump(x)
}
