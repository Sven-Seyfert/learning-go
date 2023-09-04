package main

import (
    "encoding/json"

    "github.com/davecgh/go-spew/spew"
)

type Item struct {
    ID   string                 `json:"_id"`
    Type string                 `json:"type"`
    Data map[string]interface{} `json:"-"`
}

func (i *Item) UnmarshalJSON(d []byte) error {
    var x struct {
        Item
        UnmarshalJSON struct{}
    }
    if err := json.Unmarshal(d, &x); err != nil {
        return err
    }

    var y map[string]interface{}
    _ = json.Unmarshal(d, &y)
    delete(y, "_id")
    delete(y, "type")
    *i = x.Item
    i.Data = y
    return nil
}

func main() {
	data := []byte(`[
        {"_id":"bob","type":"user","name":"Bob"},
        {"_id":"meetup","type":"website","url":"https://meetup.com/"},
        {"_id":"soup","type":"recipe","ingredients":["broth","chicken","noodles"]}
	]`)
	x := []Item{}
	if err := json.Unmarshal(data, &x); err != nil {
		panic(err)
	}
	spew.Dump(x)
}
