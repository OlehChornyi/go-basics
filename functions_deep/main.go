package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	fmt.Println(numbers)

	doubled := doubleNumbers(&numbers)
	fmt.Println(doubled)
}

func doubleNumbers(numbers *[]int) []int {
	dNumbers := []int{}
	for _, value := range *numbers {
		dNumbers = append(dNumbers, double(value))
	}

	return dNumbers
}

func double(number int) int {
	return number * 2
}
