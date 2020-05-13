package main

import (
	"fmt"
	"strings"
)

// Main method
func main() {

	// Creating and initializing string
	// Using shorthand declaration
	str1 := "!!Welcome to GeeksforGeeks !!"
	str2 := "@@This is the tutorial of Golang$$"

	// Displaying strings
	fmt.Println("Strings before trimming:")
	fmt.Println("String 1: ", str1)
	fmt.Println("String 2:", str2)

	// Trimming the given strings
	// Using Trim() function
	res1 := strings.Trim(str1, "!")
	res2 := strings.Trim(str2, "@$")

	// Trimming the given strings
	// Using TrimLeft() function
	res3 := strings.TrimLeft(str1, "!*")
	res4 := strings.TrimLeft(str2, "@")

	// Using TrimRight() function
	res5 := strings.TrimRight(str1, "!*")
	res6 := strings.TrimRight(str2, "$")

	// Trimming suffix string from the given strings
	// Using TrimSuffix() function
	res7 := strings.TrimSuffix(str1, "GeeksforGeeks")
	res8 := strings.TrimSuffix(str2, "Hello")

	// Trimming prefix string from the given strings
	// Using TrimPrefix() function
	res9 := strings.TrimPrefix(str1, "Welcome")
	res10 := strings.TrimPrefix(str2, "Hello")

	// Displaying the results
	fmt.Println("\nStrings after trimming:")
	fmt.Println("Result 1: ", res1)
	fmt.Println("Result 2:", res2)
	fmt.Println("Result 3: ", res3)
	fmt.Println("Result 4:", res4)
	fmt.Println("Result 5: ", res5)
	fmt.Println("Result 6:", res6)
	fmt.Println("Result 7: ", res7)
	fmt.Println("Result 8:", res8)
	fmt.Println("Result 7: ", res9)
	fmt.Println("Result 8:", res10)
}
