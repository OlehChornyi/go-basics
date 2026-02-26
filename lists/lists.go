package main

import "fmt"

// Slices for dynamic arrays feature
func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[1])
	fmt.Println(prices[0:1])
	prices[1] = 9.99

	updatedPrices := append(prices, 5.99)
	fmt.Println(updatedPrices)
}

// func main() {
// 	var productNames [4]string = [4]string{"A Book"}
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
// 	fmt.Println(prices)

// 	productNames[2] = "A Carpet"

// 	fmt.Println(productNames)

// 	fmt.Println(prices[2])

// 	featuredPrices := prices[1:3]
// 	pricesOne := featuredPrices[1:]
// 	pricesOne[0] = 199.99
// 	pricesTwo := pricesOne[:1]
// 	fmt.Println(pricesTwo)
// 	fmt.Println(prices)
// 	fmt.Println(len(pricesOne), cap(pricesOne))
// 	fmt.Println(pricesOne)

// 	pricesOne = pricesOne[:2]

// 	fmt.Println(len(pricesOne), cap(pricesOne))
// 	fmt.Println(pricesOne)

// 	//Modified element in a slice also mogifies an element in original array
// }
