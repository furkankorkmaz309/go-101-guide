# Composite Types in Go

## Arrays
- Fixed-length data structures.
- The length is part of the array's type.
- Can be compared using `==` and `!=`.
- Rarely used directly; mostly serve as the backing store for slices.

## Slices
- Dynamically-sized sequences.
- Use `append` to grow.
- Use `make` to set initial length and/or capacity.
- `len()` returns length, `cap()` returns capacity.
- `clear()` resets elements to zero value (length remains the same).
- `copy()` creates independent slices.
- Slices share memory—modifying one can affect others.
- Use `slices.Equal()` to compare slices (Go 1.21+).

## Strings, Bytes, and Runes
- Strings are immutable sequences of UTF-8 bytes.
- `len()` returns byte count, not character count.
- Use `[]byte` or `[]rune` conversions for manipulation.
- Beware of multi-byte characters when slicing.

## Maps
- Key-value pairs with dynamic sizing.
- Writing to a `nil` map causes a panic.
- Use `make` or `{}` to initialize.
- Use `delete()` to remove a key.
- Use `v, ok := map[key]` to safely check for key presence.
- `clear()` removes all entries.
- Use `maps.Equal()` for equality checks (Go 1.21+).
- Can be used as sets with `map[T]bool` or `map[T]struct{}`.

## Structs
- Groups related fields into one type.
- Access fields with dot notation (`.`).
- Comparable if all fields are comparable.
- Structs with the same field names, order, and types can be converted.
- Anonymous structs are useful for short-lived or ad-hoc data structures.


## 📂 Code

```go
// 01_slices_maps_structs.go
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
```
