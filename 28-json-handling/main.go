package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	createJsonFromMap()
	separator()

	createJsonFromStruct()
	separator()

	unmarshalJsonToMap()
	separator()

	unmarshalJsonToStruct()
	separator()

	apiExample()
	separator()

	apiExampleWithDecoder()
	separator()
}

func separator() {
	fmt.Println("------------------------------")
}

func createJsonFromMap() {
	// Creating JSON from a Go map.
	s := map[string]string{
		"foo": "bar",
	}

	data, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))
}

func createJsonFromStruct() {
	// Creating JSON from a Go struct.
	type person struct {
		Name        string `json:"name"`
		Age         int    `json:"age"`
		Description string `json:"descr,omitempty"`
		secret      string // Unexported fields are never (un)marshaled
	}

	s := person{
		Name:   "Bob",
		Age:    32,
		secret: "Shhh!",
	}

	data, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))
}

func unmarshalJsonToMap() {
	// Unmarshaling JSON to Go map
	data := []byte(`{"foo":"bar"}`)

	var s interface{}
	if err := json.Unmarshal(data, &s); err != nil {
		fmt.Println(err)
	}

	spew.Dump(s)
}

func unmarshalJsonToStruct() {
	// Unmarshaling JSON to Go struct
	type person struct {
		Name        string `json:"name"`
		Age         int    `json:"age"`
		Description string `json:"descr,omitempty"`
		secret      string // Unexported fields are never (un)marshaled
	}

	data := []byte(`{"name":"Bob","age":32,"secret":"Shhh!"}`)

	var s person
	if err := json.Unmarshal(data, &s); err != nil {
		fmt.Println(err)
	}

	spew.Dump(s)
}

func apiExample() {
	// Usage of unmarshal for the API response
	type todo struct {
		UserId    int    `json:"userId"`
		Id        int    `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}

	url := "https://jsonplaceholder.typicode.com/todos/1/"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal(errors.New("a status code of 'StatusOK' is expected"))
	}

	bodyData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Simple print the body as string
	fmt.Println(string(bodyData))

	// Print the body as struct
	todoItem := todo{}
	if err := json.Unmarshal(bodyData, &todoItem); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Data from API: %+v\n", todoItem)
}

func apiExampleWithDecoder() {
	// Usage of NewDecoder for the API response
	type todo struct {
		UserId    int    `json:"userId"`
		Id        int    `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"` // Comment this in or out and run the function to receive a decode error
	}

	url := "https://jsonplaceholder.typicode.com/todos/1/"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal(errors.New("a status code of 'StatusOK' is expected"))
	}

	todoItem := todo{}

	decoder := json.NewDecoder(resp.Body)

	// In case the Go struct fields and the json fields differ,
	// a error will be thrown.
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&todoItem); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Data from API: %+v\n", todoItem)
}
