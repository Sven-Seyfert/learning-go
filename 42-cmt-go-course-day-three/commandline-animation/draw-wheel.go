//nolint:go-staticcheck

package main

import (
	"fmt"
	"time"
)

func DrawWheel(s string) chan bool {
	quit := make(chan bool)

	go func() {
		for {
			for _, u := range s {
				fmt.Print(string(u))

				select {
				case <-quit:
					fmt.Print("\b \b\n")
					return
				default:
					time.Sleep(150 * time.Millisecond)
				}
				fmt.Print("\b")
			}
		}
	}()

	return quit
}

func main() {
	var q chan bool

	q = DrawWheel(`|/-\`)
	time.Sleep(5 * time.Second)

	close(q)
	time.Sleep(200 * time.Millisecond)

	q = DrawWheel(".o0OÂ°*")
	time.Sleep(5 * time.Second)
	close(q)
	time.Sleep(200 * time.Millisecond)
}
