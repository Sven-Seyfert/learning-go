package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	exampleWithWaitGroup()
	fmt.Println()

	exampleTwo()
}

func exampleWithWaitGroup() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		i := i

		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()
}

// This execution should be compared with the following
// experimental approach:
// https://github.com/golang/go/wiki/LoopvarExperiment.
func exampleTwo() {
	for i := 0; i < 100; i++ {
		n := i
		go func() {
			time.Sleep(2 * time.Millisecond)
			fmt.Println("Hallo Nr.", n)
		}()
	}
	time.Sleep(5 * time.Second)
}
