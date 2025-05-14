# Chapter 5 - Functions (Quick Notes)

## Declaring & Calling
- `func name(params) returnType {}` is the basic structure.
- Multiple inputs of the same type: `func f(x, y int)`
- Use `return` to return values; if none, omit return.

## Optional/Named Params
- No true optional/named params.
- Simulate with structs: pass a config struct with fields.

## Variadic Params
- `...type` for last param → creates a slice.
- Example: `func addAll(nums ...int)` → `nums` is a slice.

## Multiple Return Values
- Functions can return multiple values: `(int, int, error)`
- All return values must be used or ignored with `_`.

## Named Return Values
- Return vars can be named: `(x int, err error)`
- Avoid blank returns (`return` without values) — not readable.

## Functions as Values
- Functions are first-class: can be assigned, passed, returned.
- Define types with `type opFunc func(int, int) int`

## Anonymous Functions
- No name, used inline or assigned to vars.
- Common in defer and goroutines.

## Closures
- Inner functions capture and modify outer variables.
- Useful for creating stateful behaviors.

## Higher-Order Functions
- Functions can accept or return other functions.
- Used in sorting, middleware, etc.

## defer
- Defers function execution until the surrounding function ends.
- Useful for cleanup (e.g., `file.Close()`).
- LIFO order; params evaluated immediately.

## Call by Value
- All values are copied into functions.
- Maps/slices have internal pointers → changes may reflect outside.

## Tips
- Use `_` to ignore unused return values.
- Avoid blank returns and shadowing named returns.
- Use closures to reduce repeated code or manage state.


## 📂 Code

```go
// 01_functions.go
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// Return value
func div(num, denom int) float64 {
	if denom == 0 {
		return 0
	}

	// return num / denom
	// error

	floatNum := float64(num)
	floatDenom := float64(denom)

	return floatNum / floatDenom
}

// Named multiple return values
func divAndRemainder(num, denom int) (result, mod int, err error) {
	if denom == 0 {
		return 0, 0, errors.New("can not divide by 0")
	}
	result = num / denom
	mod = num % denom
	return result, mod, nil
}

// Anonymous Function and Closure
func closureExample() func() int {
	a := 5
	return func() int {
		a++
		return a
	}
}

// Use of defer
func saveResultsToFile(filename string, results []string) error {
	file, err := os.Create("05-functions/" + filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	err = encoder.Encode(results)

	if err != nil {
		return err
	}

	return nil
}

func main() {
	divide := div(73, 37)
	msg1 := fmt.Sprintf("div : %v", divide)
	msg2 := fmt.Sprintf("div : %.2f", divide)
	fmt.Println(msg1) // div : 1.972972972972973
	fmt.Println(msg2) // div : 1.97

	resultSlice := []string{msg1, msg2}
	err := saveResultsToFile("results.json", resultSlice)

	if err != nil {
		fmt.Println(err)
	}

	result, mod, err := divAndRemainder(73, 37)
	if err != nil {
		fmt.Println(err) // jump
	} else {
		fmt.Println("divAndRemainder :", result, mod) // divAndRemainder : 1 36
	}

	result, mod, err = divAndRemainder(73, 0)
	if err != nil {
		fmt.Println(err) // can not divide by 0
	} else {
		fmt.Println("divAndRemainder :", result, mod) // jump
	}

	next := closureExample()
	fmt.Println("closure 1 :", next()) // closure 1 : 6
	fmt.Println("closure 2 :", next()) // closure 2 : 7

	square := func(x int) int { return x * x }
	myNum := 3
	fmt.Printf("square of %v is %v\n", myNum, square(myNum)) // square of 3 is 9
}
```
