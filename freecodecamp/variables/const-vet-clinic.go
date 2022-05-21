package main

import (
	"fmt"
)

const (
	errorSpecialist = iota + 5
	catSpecialist
	dogSpecialist
	snakeSpecialist
)


const (
	_ = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

const (
	/*1 bit shifted by iota*/
	isAdmin = 1 << iota
	/*1 bit shifted by 1 place*/
	isHeadquarters
	/*1 bit shifted by 2 places*/
	canSeeFinancials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)


func main() {
	var specialistType int
	fmt.Printf("specialistType:  %v\n", specialistType == catSpecialist)
	fmt.Printf("catSpecialist: %v\n", catSpecialist)

	fileSize := 4200000000.
	fmt.Printf("%.2fGB\n", fileSize/GB)

	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles)
}