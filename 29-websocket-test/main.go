package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Hi from main.go.")
	fmt.Println(GetIP())
}

type IP struct {
	Query string `json:"query"`
}

func GetIP() string {
	resp, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var ip IP
	if err := json.Unmarshal(body, &ip); err != nil {
		panic(err)
	}

	return ip.Query
}
