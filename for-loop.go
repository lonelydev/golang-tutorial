package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		fmt.Printf("%v += %v equals\t", sum, i)
		sum += i
		fmt.Printf("%v\n", sum)
	}
	fmt.Println("The final sum is ", sum)
}
