package main

import "fmt"

func main() {
	basicTypes()
	aggregateTypes()
	referenceTypes()
	interfaceType()
}

// basic types (numbers, strings, booleans)
func basicTypes() {
	var myInt int // int16, int32 and int64 are almost never used
	var myUint uint
	var myFloat32 float32
	var myFloat64 float64

	myInt = 10
	myUint = 20

	myFloat32 = 10.1
	myFloat64 = 100.1

	fmt.Println(myInt, myUint, myFloat32, myFloat64)
	fmt.Printf("%d, %d, %.1f, %.1f\n", myInt, myUint, myFloat32, myFloat64)
	fmt.Println("")
}

type Car struct {
	NumberOfTires int
	Luxury        bool
	BucketSeats   bool
	Make          string
	Model         string
	Year          int
}

// aggregate types (array, struct)
func aggregateTypes() {
	// array
	var myStrings [3]string

	myStrings[0] = "cat"
	myStrings[1] = "dog"
	myStrings[2] = "fish"

	fmt.Println("First element is", myStrings[0])
	fmt.Println("")

	// struct
	var myFirstCar Car
	myFirstCar.NumberOfTires = 4
	myFirstCar.Luxury = false
	myFirstCar.BucketSeats = false
	myFirstCar.Make = "Volkswagen"
	myFirstCar.Model = "Polo"
	myFirstCar.Year = 2004

	fmt.Println("First car year:", myFirstCar.Year)

	mySecondCar := Car{
		NumberOfTires: 4,
		Luxury:        true,
		BucketSeats:   false,
		Make:          "Porsche",
		Model:         "911",
		Year:          2013,
	}

	fmt.Printf("Second car make is %s. The number of tires are %d.", mySecondCar.Make, myFirstCar.NumberOfTires)
}

// reference types (pointers, slices, maps, functions, channels)
func referenceTypes() {

}

// interface type
func interfaceType() {

}
