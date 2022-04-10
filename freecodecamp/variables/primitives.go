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

	// unsigned integer
	var unSignedInt uint16 = 42
	fmt.Printf("the value of unSignedInt is %v and type is %T\n", unSignedInt, unSignedInt)

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

	fmt.Printf("\n\n*****+*****+*****+*****\n")
	fmt.Printf("Let us check out logical operations\n")
	fmt.Printf("*****+*****+*****+*****\n")

	// logical operations on int
	orResult := a | b // 1011
	fmt.Printf("the value of orResult is %v and type is %T\n", orResult, orResult)
	andResult := a & b // 0010
	fmt.Printf("the value of andResult is %v and type is %T\n", andResult, andResult)
	xorResult := a ^ b // 1001 
	fmt.Printf("the value of xorResult is %v and type is %T\n", xorResult, xorResult)
	nandResult := a &^ b // 0100
	fmt.Printf("the value of nandResult is %v and type is %T\n", nandResult, nandResult)
	
	// float 
	fmt.Printf("\n\n*****+*****+*****+*****\n")
	fmt.Printf("Let us check out floats\n")
	fmt.Printf("*****+*****+*****+*****\n")
	floatN := 3.14
	floatN = 13.7e72
	floatN = 2.1E14
	fmt.Printf("the value of floatN is %v and type is %T\n", floatN, floatN)

	fmt.Printf("\n\n*****+*****+*****+*****\n")
	fmt.Printf("Let us check out complex numbers\n")
	fmt.Printf("*****+*****+*****+*****\n")
	var complexNumber complex64 = 1+2i
	fmt.Printf("the value of complexNumber is %v and type is %T\n", complexNumber, complexNumber)
	fmt.Printf("the value of real(complexNumber1) is %v and type is %T\n", real(complexNumber), real(complexNumber))
	fmt.Printf("the value of imag(complexNumber1) is %v and type is %T\n", imag(complexNumber), imag(complexNumber))

	// how to make a complex number?
	var complexNumber1 complex128 = complex(5, 12)
	fmt.Printf("the value of complexNumber1 is %v and type is %T\n", complexNumber1, complexNumber1)
	fmt.Printf("the value of real(complexNumber1) is %v and type is %T\n", real(complexNumber1), real(complexNumber1))
	fmt.Printf("the value of imag(complexNumber1) is %v and type is %T\n", imag(complexNumber1), imag(complexNumber1))

	fmt.Printf("\n\n*****+*****+*****+*****\n")
	fmt.Printf("Let us check out strings\n")
	fmt.Printf("*****+*****+*****+*****\n")

	stringExplore := "this is a string"
	fmt.Printf("the value of stringExplore is '%v' and type is %T\n", stringExplore, stringExplore)
	fmt.Printf("the value of stringExplore[2] is '%v' and type is %T\n", stringExplore[2], stringExplore[2])
	fmt.Printf("the value of string(stringExplore[2]) is '%v' and type is %T\n", string(stringExplore[2]), string(stringExplore[2]))

}