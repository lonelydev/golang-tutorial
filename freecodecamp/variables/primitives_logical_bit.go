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

	fmt.Printf("\n\n*****+*****+*****+*****\n")
	fmt.Printf("Let us do some bit shifting operations\n")
	fmt.Printf("*****+*****+*****+*****\n")

	// 01, 10, 11, 100, 101, 110, 111, 1000
	// 00001000
	// 00010000
	// 00100000
	// 01000000
	bitShiftVar := 8
	fmt.Printf("the value of bitShiftVar is %v\n", bitShiftVar)
	fmt.Printf("the value of bitShiftVar left shifted once is %v\n", bitShiftVar << 1)
	fmt.Printf("the value of bitShiftVar left shifted three times is %v\n", bitShiftVar << 3)
	fmt.Printf("the value of bitShiftVar right shifted once is %v\n", bitShiftVar >> 1)
	fmt.Printf("the value of bitShiftVar right shifted twice is %v\n", bitShiftVar >> 2)
}