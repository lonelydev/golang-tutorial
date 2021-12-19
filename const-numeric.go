package main

import "fmt"

const (
	/*
		Creating a huge number by shifting a bit left 100 places.
		binary 1 followed by a 100 zeroes.
	*/
	Big = 1 << 100
	/*
		Shift Big right by 99 places
	*/
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(Small)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(Big)
	fmt.Println(needFloat(Big))
}
