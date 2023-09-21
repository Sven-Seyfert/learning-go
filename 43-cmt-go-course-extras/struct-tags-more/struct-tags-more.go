package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name  string `mytag:"MyName" other:"N"`
	Email string `mytag:"MyEmail" other:"M"`
}

func main() {
	u := User{"Bob", "bob@mycompany.com"}
	t := reflect.TypeOf(u)

	fmt.Printf("Tag: %v (%T)\n", t, t)

	for _, fieldName := range []string{"Name", "Email"} {
		field, found := t.FieldByName(fieldName)
		if !found {
			continue
		}

		fmt.Printf("Field: User.%s  [%v (%T)]\n", fieldName, field, field)
		fmt.Printf("\tWhole tag value : %q\n", field.Tag)
		fmt.Printf("\tValue of 'mytag': %q\n", field.Tag.Get("mytag"))
		fmt.Printf("\tValue of 'mytag': %q\n", field.Tag.Get("other"))
	}
}
