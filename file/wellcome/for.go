package main

import {
	"fmt"
	"math"
}
	

func main() {
	c := Circle{0, 0, 5}
	r := Rectangle{0, 0, 10, 10}

	fmt.Println(totalArea(&c, &r))
}
