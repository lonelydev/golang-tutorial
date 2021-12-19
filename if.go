package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
func main() {
	var x float64 = 112.543
	var y float64 = -112.543
	fmt.Println(sqrt(x))
	fmt.Println(sqrt(y))
}
