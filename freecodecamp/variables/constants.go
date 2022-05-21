package main

import (
	"fmt"
)

// like variables you can work with constants in a block
// iota is a keyword just says default
// enumerated constants are declared as a block
// do this at a package level
const (
	err = iota
	a
	b
	c
)

func main() {
	var d int16 = 27
	fmt.Printf("value of a is %v, and type is %T\n", a, a)	
	fmt.Printf("value of a, b, c, d are is %v, %v, %v, %v\n", a, b, c, d)
	fmt.Printf("value of a + b is %v, and type is %T\n", a + b, a + b)
	fmt.Printf("value of a + d is %v, and type is %T\n", a + d, a + d)
}