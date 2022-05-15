package main

import (
	"fmt"
)

func main() {
	fmt.Printf("\n\n*****+*****+*****+*****\n")
	fmt.Printf("Let us explore strings and runes\n")
	fmt.Printf("*****+*****+*****+*****\n")

	stringExplore := "eakan"
	fmt.Printf("the value of stringExplore is '%v' and type is %T\n", stringExplore, stringExplore)
	fmt.Printf("the value of stringExplore[2] is '%v' and type is %T\n", stringExplore[2], stringExplore[2])
	fmt.Printf("the value of string(stringExplore[2]) is '%v' and type is %T\n", string(stringExplore[2]), string(stringExplore[2]))

	byteString := []byte(stringExplore)
	fmt.Printf("byteString is: %v,\nIt is of type: %T\n", byteString, byteString)

	simpleRune := 'e'
	fmt.Printf("simpleRune is: %v,\nIt is of type: %T\n", simpleRune, simpleRune)

	// runes are aliased to int32
	// bytes are int16
	// strings are made of bytes and they can contain valid characters
	// that can be represented by runes
	// rune() function  converts string to an array of runes
	runeString := []rune(stringExplore)
	fmt.Printf("runeString is: %v,\nIt is of type: %T\n", runeString, runeString)

}