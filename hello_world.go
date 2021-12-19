package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello, Eakan here!")
	fmt.Println("The time is", time.Now())
	fmt.Println("My favourite number is", rand.Intn(90))
	fmt.Println(add(2, 3))
	fmt.Println(swap("what", "now"))

	var c, python, java = false, true, false
	var i, j int = 2, 3
	fmt.Println(i, j, c, python, java)
	k := 3
	fmt.Println("k is", k)
	l, m, n := true, true, false
	fmt.Println(i, j, c, python, java, l, m, n)
}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (first, last string) {
	first = y
	last = x
	return // naked return
}
