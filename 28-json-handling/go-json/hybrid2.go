package main

import (
    "encoding/json"
    "fmt"
)

type Item struct {
    ID   string                 `json:"_id"`
    Type string                 `json:"type"`
    Data map[string]interface{} `json:"-"`
}

func (i *Item) MarshalJSON() ([]byte, error) {
    data, _ := json.Marshal(i.Data)
    tmp := struct{
        *Item
        MarshalJSON struct{} `json:"-"`
    }{Item: i}
    obj, _ := json.Marshal(tmp)
    obj[len(obj)-1] = ','
    return append(obj, data[1:]...), nil
}

func main() {
    x := []Item{
        {ID: "bob", Type:"user",Data:map[string]interface{}{"name":"Bob"}},
        {ID:"meetup",Type:"website",Data:map[string]interface{}{"url":"https://meetup.com/"}},
        {ID:"soup",Type:"recipe",Data:map[string]interface{}{"ingredients":[]string{"broth","chicken","noodles"}}},
    }
    data, err := json.MarshalIndent(x, "", "  ")
    if err != nil {
        panic(err)
    }
    fmt.Println(string(data))
}
