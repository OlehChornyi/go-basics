package anonymous

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3}

	double := createTransformer(2)
	triple := createTransformer(3)

	//Anonymous function
	transformed := transformNumbers(&numbers, func(number int) int { return number * 2 })

	doubled := transformNumbers(&numbers, double)
	fmt.Println(doubled)

	tripled := transformNumbers(&numbers, triple)
	fmt.Println(tripled)

	fmt.Println(transformed)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

// Closure
func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

// type transformFn func(int) int

// // Functions as firs-class values (pass functions)
// func main() {
// 	numbers := []int{1, 2, 3, 4}
// 	fmt.Println(numbers)

// 	doubled := transformNumbers(&numbers, double)
// 	fmt.Println(doubled)

// 	tripled := transformNumbers(&numbers, triple)
// 	fmt.Println(tripled)

// 	moreNumbers := []int{5, 1, 2}
// 	transformerFn1 := getTransformerFunction(&numbers)
// 	transformerFn2 := getTransformerFunction(&moreNumbers)

// 	transformedNumbers := transformNumbers(&numbers, transformerFn1)
// 	moreTransformedNumbers := transformNumbers(&moreNumbers, transformerFn2)

// 	fmt.Println(transformedNumbers)
// 	fmt.Println(moreTransformedNumbers)

// }

// func transformNumbers(numbers *[]int, transform transformFn) []int {
// 	dNumbers := []int{}
// 	for _, value := range *numbers {
// 		dNumbers = append(dNumbers, transform(value))
// 	}

// 	return dNumbers
// }

// func getTransformerFunction(numbers *[]int) transformFn {
// 	if (*numbers)[0] == 1 {
// 		return double
// 	} else {
// 		return triple
// 	}
// }

// func double(number int) int {
// 	return number * 2
// }

// func triple(number int) int {
// 	return number * 3
// }
