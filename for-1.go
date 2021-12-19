package main

import "fmt"

func main() {
	sum := 1
	for sum < 10 {
		fmt.Printf("%v += %v equals\t", sum, sum)
		sum += sum
		fmt.Printf("%v\n", sum)
	}
	fmt.Println("The final sum is ", sum)
}
