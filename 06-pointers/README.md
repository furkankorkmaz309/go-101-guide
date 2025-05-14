# Chapter 6 - Pointers (Quick Notes)

## Basics
- A pointer holds a memory address of a value.
- `&` gets the address, `*` gets/sets the value (dereference).
- Nil is the zero value for pointers.
- You can’t take the address of a constant or literal.

## new vs &T{}
- `new(T)` allocates zero value of T and returns pointer.
- Prefer `&T{}` for structs.
- For primitive literals, create variable or use helper: `makePointer("val")`.

## Function Behavior
- Go is call-by-value: function gets a copy.
- Copying a pointer = access to same data.
- To update original data → dereference inside the function (`*ptr = value`).
- Reassigning pointer in func doesn’t affect original.

## Pointers & Mutability
- Use pointers to allow mutation in functions.
- Value parameters = safe, immutable behavior by default.
- Avoid pointer unless mutation or performance is needed.

## Pointers Are Rarely Needed
- Instead of passing a pointer to fill a struct, return the struct.
- JSON is an exception: needs pointer to fill data into passed-in variable.

## Performance Tips
- Return values unless struct is very large (e.g. >10MB).
- Returning small pointers can be slower than returning values.
- For tight loops or large structs, pointers can reduce GC pressure.

## Zero Value vs No Value
- Use nil pointer to signal “no value” (e.g. `*string = nil`).
- But be cautious — only do this if truly needed (e.g., in JSON decoding).

## Maps & Slices
- Maps are pointers internally → all changes are reflected.
- Slices: content changes are visible, length changes aren't.
- Appending may break reference due to new allocation.

## Buffers
- Use pre-allocated slices as reusable buffers to reduce memory garbage.
- Avoid allocating inside tight loops (e.g., file reading).

## Garbage Collector
- Stack is fast and limited, heap is managed by GC and slower.
- Prefer stack-allocated values (known size, not returned).
- Returning pointers forces heap allocation (escape analysis).
- Set `GOGC` to control GC frequency, `GOMEMLIMIT` to cap memory.

## Summary
- Use values unless mutation or performance requires pointers.
- Avoid unnecessary allocations.
- Document mutation clearly when using pointer params.


## 📂 Code

```go
// 01_pointers.go
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
```
