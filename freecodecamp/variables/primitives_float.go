package main

import (
	"fmt"
)

func main() {
	// float 
	fmt.Printf("\n\n*****+*****+*****+*****\n")
	fmt.Printf("Let us check out floats\n")
	fmt.Printf("*****+*****+*****+*****\n")
	floatN := 3.14
	floatN = 13.7e72
	floatN = 2.1E14
	fmt.Printf("the value of floatN is %v and type is %T\n", floatN, floatN)
}