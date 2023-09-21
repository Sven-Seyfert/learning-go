package main

import (
	"fmt"
	"math"
)

type AvgFunc func(float64, float64) float64

func compute(fn AvgFunc) float64 {
	return fn(3, 4)
}

func average(x, y float64) float64 {
	return (x + y) / 2
}

func geommean(x, y float64) float64 {
	return math.Sqrt(x * y)
}

func harmmean(x, y float64) float64 {
	return 1.0 / (1.0/x + 1.0/y)
}

func main() {
	fmt.Println(compute(average))

	avg := average
	fmt.Println(compute(avg))

	avg = geommean
	fmt.Println(compute(avg))

	avg = harmmean
	fmt.Println(compute(avg))
	fmt.Println(compute(math.Pow))

	//  --------------

	f := func() {
		fmt.Println("Hello, here is an anonymous function!")
	}

	fmt.Printf("(%T)\n", f)

	f()
}
