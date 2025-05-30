package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("Array")
	var myArray1 = [3]int{1, 2, 3}
	var myArray2 = [...]int{1, 2, 3}
	fmt.Println(myArray1 == myArray2) // true

	myArray1[1] = 9
	fmt.Println(myArray1 == myArray2) // false

	// It won't work!
	// fmt.Println(slices.Equal(myArray1, myArray2))
	mySliceFromArray1 := myArray1[:]
	mySliceFromArray2 := myArray2[:]
	fmt.Println(slices.Equal(mySliceFromArray1, mySliceFromArray2)) // false

	fmt.Println("\nSlice")
	mySlice := []int{10, 20, 30}
	fmt.Println(mySlice, len(mySlice), cap(mySlice)) // [10 20 30] 3 3

	mySlice = append(mySlice, 40, 50)
	fmt.Println(mySlice, len(mySlice), cap(mySlice)) // [10 20 30 40 50] 5 6

	subSlice := mySlice[:3]
	fmt.Println(subSlice, len(subSlice), cap(subSlice)) // [10 20 30] 3 6

	newSlice := make([]int, 3)
	copy(newSlice, mySlice[1:])
	fmt.Println(newSlice, len(newSlice), cap(newSlice)) // [20 30 40] 3 3

	clear(subSlice)
	fmt.Println(subSlice, len(subSlice), cap(subSlice)) // [0 0 0] 3 6

	x := []string{"a", "b", "c", "d"}
	y := x[:2]
	z := x[1:]

	x[1] = "y"
	y[0] = "x"
	z[1] = "z"

	fmt.Println("x :", x) // x : [x y z d]
	fmt.Println("y :", y) // y : [x y]
	fmt.Println("z :", z) // z : [y z d]

	fmt.Println("\nMap")
	myMap := map[string]int{"Alice": 5}
	myMap["Bob"] = 3
	fmt.Println(myMap) // map[Alice:5 Bob:3]

	delete(myMap, "Alice")
	fmt.Println(myMap) // map[Bob:3]

	v, ok := myMap["Bob"]
	fmt.Println(v, ok) // 3 true

	v, ok = myMap["Alice"]
	fmt.Println(v, ok) // 0 false

	myMap["Bob"]++
	fmt.Println(myMap) // map[Bob:4]

	fmt.Println("\nStruct")
	type Person struct {
		Name string
		Age  int
	}
	person1 := Person{"Alice", 30}
	person2 := Person{Name: "Bob", Age: 40}
	var person3 Person
	person3.Name = "Charlie"
	person3.Age = 50
	fmt.Println(person1, person2, person3) // {Alice 30} {Bob 40} {Charlie 50}
}
