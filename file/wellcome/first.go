// Go program to illustrate the concept
// of dynamic values and types
package main

import (
	"fmt"
)

// Creating an interface
func myInterface(a interface{}) {
	value, ok := a.(float64)
	fmt.Println(value, ok)
}
func main() {
	var a1 interface {
	} = 98.09

	myInterface(a1)

	var a2 interface {
	} = "Germa 66"

	myInterface(a2)
}
