package main

import "fmt"

func main() {
	// Pointer 101
	myNum := 73
	myPointer := &myNum
	fmt.Println("Address of myNum is", myPointer) // Address of myNum is 0xc000010100
	fmt.Println("Value of myPointer is", myNum)   // Value of myPointer is 73

	*myPointer = 20
	fmt.Println("Updated myNum is", myNum) // Updated myNum is 20

	// nil pointer
	var myPointer2 *int
	fmt.Println("Is myPointer 2 nil?", myPointer2 == nil) // Is myPointer 2 nil? true
	// fmt.Println(*myPointer2)
	// Panic! Because of dereference

	// Updating pointer
	myNum2 := 37
	failedUpdate := func(ptr *int) {
		x := 99
		ptr = &x
	}

	successfulUpdate := func(ptr *int) {
		*ptr = 99
	}

	failedUpdate(&myNum2)
	fmt.Println("Value after failedUpdate :", myNum2) // Value after failedUpdate : 37

	successfulUpdate(&myNum2)
	fmt.Println("Value after successfulUpdate :", myNum2) // Value after successfulUpdate : 99
}
