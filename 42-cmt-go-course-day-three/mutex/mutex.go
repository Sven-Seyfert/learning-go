package main

import (
	"fmt"
	"time"
)

type SafeCounter struct {
	v int
	// mux sync.Mutex
}

func (c *SafeCounter) Inc() {
	// c.mux.Lock()
	// defer c.mux.Unlock()
	c.v++
}

func (c *SafeCounter) Value() int {
	// c.mux.Lock()
	// defer c.mux.Unlock()
	return c.v
}

// Mutex example which should be used with care.
func main() {
	for {
		c := SafeCounter{}
		for i := 0; i < 100000; i++ {
			go c.Inc()
		}
		time.Sleep(time.Second)
		fmt.Println(c.Value())
	}
}
