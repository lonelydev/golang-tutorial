package main

import (
	"fmt"
)

func main() {
	fmt.Printf("\n\n*****+*****+*****+*****\n")
	fmt.Printf("Let us check out strings\n")
	fmt.Printf("*****+*****+*****+*****\n")

	stringExplore := "this is a string"
	fmt.Printf("the value of stringExplore is '%v' and type is %T\n", stringExplore, stringExplore)
	fmt.Printf("the value of stringExplore[2] is '%v' and type is %T\n", stringExplore[2], stringExplore[2])
	fmt.Printf("the value of string(stringExplore[2]) is '%v' and type is %T\n", string(stringExplore[2]), string(stringExplore[2]))

}