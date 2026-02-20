package main

import "fmt"

func main() {
	age := 32

	var agePointer *int
	agePointer = &age

	fmt.Println("Age:", *agePointer)
	fmt.Println(getAdultYears(agePointer))
}

func getAdultYears(age *int) int {
	return *age - 18
}
