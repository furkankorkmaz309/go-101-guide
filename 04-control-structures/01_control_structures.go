package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Blocks
	myNum1 := 10
	if myNum1 > 5 {
		fmt.Println(myNum1) // 10
		myNum1 := 5
		fmt.Println(myNum1) // 5
	}
	fmt.Println(myNum1) // 10

	if myNum2 := rand.Intn(10); myNum2 == 0 {
		fmt.Println("That's too low")
	} else if myNum2 > 5 {
		fmt.Println("That's too big :", myNum2)
	} else {
		fmt.Println("That's a good number :", myNum2)
	}
	// fmt.Println(myNum2)
	// undefined error

	// For loops
	fmt.Print("(")
	for i := 0; i < 10; i++ {
		fmt.Print(i, ",")
	}
	fmt.Print("10)\n") // (0,1,2,3,4,5,6,7,8,9,10)

	i := 0
	for {
		fmt.Println("Try")
		i++
		if i == 3 {
			break
		}
	}

	mySlice := []int{1, 2, 3, 4, 5}

	for i, v := range mySlice {
		fmt.Printf("Index : %v | Value : %v\n", i, v)
	}

	// switch case
	for i := 11; i > 0; i-- {
		var modulus int = i % 5

		switch modulus {
		case 0:
			fmt.Printf("%v Nice!\n", i)
		case 1:
			fmt.Printf("%v So close!\n", i)
		default:
			fmt.Printf("%v %v 5 = %v\n", i, "%", modulus)
		}
	}
}
