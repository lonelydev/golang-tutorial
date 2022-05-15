package main

import (
	"fmt"
)

func main() {
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
}