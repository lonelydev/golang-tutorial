package main

import (
	"fmt"
)

func main() {
	var students [3]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "Lisa"
	students[1] = "Ahmed"
	students[2] = "David"
	fmt.Printf("Students: %v\n", students)

	var grades = [3]int{5, 9, 13}
	fmt.Printf("Grades: %v\n", grades)
	var gradesBetterSyntax = [...]int{0, 9, 13}
	fmt.Printf("GradesBetterSyntax: %v\n", gradesBetterSyntax)
	fmt.Printf("size of GradesBetterSyntax: %v\n", len(gradesBetterSyntax))
	fmt.Printf("GradesBetterSyntax[1]: %v\n", gradesBetterSyntax[1])

	// linear algebra
	// array of arrays - matrix
	var identityMatrix [3][3]int = [3][3]int{ {1, 0, 0}, {0,1,0}, {0,0,1}}
	fmt.Printf("identityMatrix: %v\n", identityMatrix)

	// in golang - arrays are values!
	// when you copy an array, you are creating a literal copy and not a pointer change
	// this means copying an array can be expensive - keep that in mind
	array1 := [...]int{1,3, 5}
	array2 := array1
	array2[2] = 7
	fmt.Printf("array1: %v\n", array1)
	fmt.Printf("array2: %v\n", array2)

	// in order to avoid the cost involved in large copy operations
	// you could use pointers using address references.
	array3 := [...]int{1,3, 5}
	array4 := &array3
	array4[2] = 7
	fmt.Printf("array1: %v\n", array3)
	fmt.Printf("array2: %v\n", array4)

	// introducing slices - they behave more like traditional arrays in other languages
	// slices are initialised by initialising an array without a size.
	slice1 := []int{1, 2, 3}
	slice2 := slice1
	fmt.Printf("slice1: %v\n", slice1)
	fmt.Printf("slice2: %v\n", slice2)
	slice2[1] = 33
	fmt.Printf("slice2 after modification: %v\n", slice2)
	fmt.Printf("slice1 after slice2 modification: %v\n", slice1)
	fmt.Printf("len(slice1): %v\n", len(slice1))
	fmt.Printf("cap(slice1): %v\n", cap(slice1))

	// other cool stuff you can do with a slice
	oneToTenSlice := []int{1,2,3,4,5,6,7,8,9,0}
	fmt.Printf("oneToTenSlice: %v\n", oneToTenSlice)
	// slice of all elements
	secondSlice := oneToTenSlice[:]
	fmt.Printf("secondSlice: %v\n", secondSlice)
	
	// slice of all elements from the 4th element
	thirdSlice := oneToTenSlice[3:]
	fmt.Printf("thirdSlice: %v\n", thirdSlice)

	// slice of all elements upto the 4th element
	fourthSlice := oneToTenSlice[:3]
	fmt.Printf("fourthSlice: %v\n", fourthSlice)

	fmt.Printf("original oneToTenSlice: %v\n", oneToTenSlice)

	testSlice := []int{1,2,3,4,5}
	fmt.Printf("testSlice: %v\n", testSlice)

	lastFourTestSlice := testSlice[1:]
	fmt.Printf("lastFourTestSlice: %v\n", lastFourTestSlice)

	firstFourTestSlice := testSlice[:len(testSlice)-1]
	fmt.Printf("firstFourTestSlice: %v\n", firstFourTestSlice)

}