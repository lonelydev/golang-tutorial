package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	delta := z*z - x
	i := 1
	prevDelta := 0.0
	for math.Abs(delta) > 0 && delta != prevDelta*-1 {
		z -= delta / (2 * z)
		fmt.Printf("Attempt %v, delta is %v and z is %v\n", i, delta, z)
		i += 1
		prevDelta = delta
		delta = z*z - x
	}
	return z
}

func main() {
	fmt.Println(Sqrt(625))
}
