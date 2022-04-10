package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int
	i = 42
	var j int = 27
	k := 99

	var actorName string = "Eakan Gopalakrishnan"
	var companion string = "Durga Krishnan"

	var (
		actorName1 string = "Ken Gokrish"
		companion1 string = "Dur Krish"
		anotherInt int = 43
		stringAnotherIntWrong string = string(anotherInt)
		stringAnotherIntRight string = strconv.Itoa(anotherInt)
		firstFloat float32 = 43.012345
	)

	fmt.Printf("Value of i is %v and type is %T \n",i, i)
	fmt.Printf("Value of j is %v and type is %T \n",j, j)
	fmt.Printf("Value of k is %v and type is %T \n",k, k)
	fmt.Printf("Value of actorName is %v and type is %T \n", actorName, actorName)
	fmt.Printf("Value of companion is %v and type is %T \n", companion, companion)
	fmt.Printf("Value of actorName1 is %v and type is %T \n", actorName1, actorName1)
	fmt.Printf("Value of companion1 is %v and type is %T \n", companion1, companion1)
	fmt.Printf("Value of anotherInt is %v and type is %T \n", anotherInt, anotherInt)
	fmt.Printf("Value of firstFloat is %v and type is %T \n", firstFloat, firstFloat)
	fmt.Printf("Value of stringAnotherIntWrong is %v and type is %T \n", stringAnotherIntWrong, stringAnotherIntWrong)
	fmt.Printf("Value of stringAnotherIntRight is %v and type is %T \n", stringAnotherIntRight, stringAnotherIntRight)
}