# Notes Predeclared Types and Declarations

## Boolean Type

* In Go, the type `bool` can only take two values: `true` or `false`.
* The default (zero) value is `false`.
* Example: `var myBool bool = true`

## Numeric Types

### Signed and Unsigned Integers

* The `int` (signed) and `uint` (unsigned) variable types are available in different sizes.
* `byte`: is an alias for `uint8`.
* `rune`: Alias for `int32`, representing a single Unicode character.

#### Notes:

* Defaults to `0`.
* Fixed-size types (`int32`, `uint64`, etc.) are used especially for network protocols or file formats.
* It is recommended to use `int` unless otherwise required.

#### Decimal Numbers (Floating Point)

* There are two float types in Go: `float32` and `float64`.
* `float64` is preferred because it is more precise.
* The default value is `0.0`.
* Example: `var euler float32 = 2.718`

### Constants

* Defined with `const`; its value is fixed and cannot be changed.
* Must be computable at compile time.
* Can be used for numbers, characters and strings.
* Example: `const Pi = 3.14`

### Strings

* The `string` type represents UTF-8 character strings.
* It is an immutable data type.
* Default value: `""` (empty string)
* Example: `var hello string = "Hello, Go!"`

### Zero Value (Default Values)

* In Go, if a variable is defined but not assigned a value, the language automatically assigns a "zero value":

  * `int` → `0`
  * `bool` → `false`
  * `string` → `""` (empty)

### Short Declaration

* With `:=`, a variable can be both defined and assigned a value only in functions.
* Example: `x := 10`

### Type Conversion

* There is no automatic type conversion, each conversion must be done explicitly.
* It is done with conversions like `float64(myInt)`.
* This helps to make it clear what the code is doing.

#### Extra Information

* The `reflect.TypeOf(...)` function allows you to print the type of a variable in the program.
* Unused local variables will generate a compilation error (but not constants).

Translated with DeepL.com (free version)

## 📂 Code

```go
// 01_types.go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Boolean
	var myBool bool = true
	fmt.Println("Boolean:", myBool)

	// Integers
	var myInt int = 42
	var myUint uint = 100
	var myByte byte = 255 // same as uint8
	var myRune rune = 'G' // same as int32
	fmt.Println("Integers:", myInt, myUint, myByte, myRune)

	// Floating Points
	var myNum float64 = 10.0
	var euler float32 = 2.718
	fmt.Println("Floats:", myNum, euler)

	// Strings
	var hello string = "Hello, Go!"
	fmt.Println("String:", hello)

	// Zero values
	var defaultInt int
	var defaultBool bool
	var defaultString string
	fmt.Println("Zero values:", defaultInt, defaultBool, defaultString)

	// Constants
	const Pi = 3.14
	const Language = "Go"
	fmt.Println("Constants:", Pi, Language)

	// Short declaration
	x := 10
	y := "short declaration"
	fmt.Println("Short declaration:", x, y)

	// Type conversion
	myInt = 5
	var myIntConverted float64 = float64(myInt)
	fmt.Printf("\nType conversion\n%v (%v)\n%v (%v)\n", myInt, reflect.TypeOf(myInt), myIntConverted, reflect.TypeOf(myIntConverted))
}
```
