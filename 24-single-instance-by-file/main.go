package main

import (
	"24-single-instance-by-file/singleinstance"

	"fmt"
	"time"
)

func main() {
	// check for second instance
	// if it's the first, create a instance file as indicator
	singleinstance.CreateInstanceFile()

	// do something
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}

	// this is temporary to test multiple calls of this programm
	time.Sleep(2 * time.Second)

	// remove the previously created instance file
	defer singleinstance.RemoveInstanceFile()
}
