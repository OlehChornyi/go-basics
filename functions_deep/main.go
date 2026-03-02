package main

import "fmt"

// Variadic functions work with any ammount of parameters
func main() {
	// numbers := []int{1, 10, 15}
	sum := sumup(1, 10, 15, 40, -5)

	fmt.Println(sum)
}

func sumup(startingValue int, numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}
