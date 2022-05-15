package main

import (
	"fmt"
)

func main() {
	var isSomething bool = false
	var isTrue bool = true
	var isNotInitialised bool

	fmt.Printf("\n\n*****+*****+*****+*****\n")
	fmt.Printf("Let us check out booleans\n")
	fmt.Printf("*****+*****+*****+*****\n")

	fmt.Printf("the value of isSomething is %v and type is %T\n", isSomething, isSomething)
	fmt.Printf("the value of isTrue is %v and type is %T\n", isTrue, isTrue)
	fmt.Printf("the value of isNotInitialised is %v and type is %T\n", isNotInitialised, isNotInitialised)
	// big package can help working with very large numbers

	fmt.Printf("\n\n*****+*****+*****+*****\n")
	fmt.Printf("Let us check out mathematical operations\n")
	fmt.Printf("*****+*****+*****+*****\n")

	// arithmetic operations on int
	var a int = 10 // 1010
	var b int = 3 // 0011
	c := a / b
	fmt.Printf("the value of a is %v and type is %T\n", a, a)
	fmt.Printf("the value of b is %v and type is %T\n", b, b)
	fmt.Printf("the value of c is %v and type is %T\n", c, c)

}