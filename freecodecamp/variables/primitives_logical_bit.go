package main

import (
	"fmt"
)

func main() {
	// arithmetic operations on int
	var a int = 10 // 1010
	var b int = 3 // 0011
	
	fmt.Printf("\n\n*****+*****+*****+*****\n")
	fmt.Printf("Let us check out logical operations\n")
	fmt.Printf("*****+*****+*****+*****\n")

	fmt.Printf("a is %v and b is %v\n", a, b)
	// logical operations on int
	orResult := a | b // 1011
	fmt.Printf("the value of a|b is %v and type is %T\n", orResult, orResult)
	andResult := a & b // 0010
	fmt.Printf("the value of a&b is %v and type is %T\n", andResult, andResult)
	xorResult := a ^ b // 1001 
	fmt.Printf("the value of a^b (a xor b) is %v and type is %T\n", xorResult, xorResult)
	nandResult := a &^ b // 0100
	fmt.Printf("the value of a&^b (a nand b) is %v and type is %T\n", nandResult, nandResult)
}