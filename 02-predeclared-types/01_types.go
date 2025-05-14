package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Boolean
	var myBool bool = true
	fmt.Println("Boolean:", myBool) // Boolean: true

	// Integers
	var myInt int = 42
	var myUint uint = 100
	var myByte byte = 255                                   // same as uint8
	var myRune rune = 'G'                                   // same as int32
	fmt.Println("Integers:", myInt, myUint, myByte, myRune) // Integers: 42 100 255 71

	// Floating Points
	var myNum float64 = 10.0
	var euler float32 = 2.718
	fmt.Println("Floats:", myNum, euler) // Floats: 10 2.718

	// Strings
	var hello string = "Hello, Go!"
	fmt.Println("String:", hello) // String: Hello, Go!

	// Zero values
	var defaultInt int
	var defaultBool bool
	var defaultString string
	fmt.Printf("Zero values\nInt : %v\nBool : %v\nString : %v\n", defaultInt, defaultBool, defaultString)
	/*
		Zero values
		Int : 0
		Bool : false
		String :
	*/

	// Constants
	const Pi = 3.14
	const Language = "Go"
	fmt.Println("Constants:", Pi, Language) // Constants: 3.14 Go

	// Short declaration
	x := 10
	y := "short declaration"
	fmt.Println("Short declaration:", x, y) // Short declaration: 10 short declaration

	// Type conversion
	myInt = 5
	var myIntConverted float64 = float64(myInt)
	fmt.Printf("\nType conversion\n%v (%v)\n%v (%v)\n", myInt, reflect.TypeOf(myInt), myIntConverted, reflect.TypeOf(myIntConverted))
	/*
		Type conversion
		5 (int)
		5 (float64)
	*/
}
