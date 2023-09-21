package main

import (
	"fmt"
	"time"
)

const itWorks = true

func provide() {
	if itWorks {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Fehler:", r)
			}
		}()
	}
	panic("Autsch!")
}

// A panic within a goroutine cannot be catched outside of
// the goroutine (in a other or the main routine).
func main() {
	if !itWorks {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Fehler:", r)
			}
		}()
	}

	go provide()
	time.Sleep(5 * time.Second)
}
