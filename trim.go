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

	// Displaying the results
	fmt.Println("\nStrings after trimming:")
	fmt.Println("Result 1: ", res1)
	fmt.Println("Result 2:", res2)
	fmt.Println("Result 3: ", res3)
	fmt.Println("Result 4:", res4)
	fmt.Println("Result 5: ", res5)
	fmt.Println("Result 6:", res6)
}
