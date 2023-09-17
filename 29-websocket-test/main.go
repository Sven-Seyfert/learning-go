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
	Query string
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
	json.Unmarshal(body, &ip)

	return ip.Query
}
