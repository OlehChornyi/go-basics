package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	years, expectedReturnRate := 10.0, 5.5

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Investment Horizon(years): ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for infliation): %.2f\n", futureRealValue)
	// fmt.Println("Future Value:",futureValue)
	// fmt.Printf("Future Value: %.2f\nFuture Value (adjusted for infliation): %.2f", futureValue, futureRealValue)
	// fmt.Println("Future Value (adjusted for infliation):", futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
}

// Profit Calculator
// func main() {
// 	var revenue float64
// 	var expences float64
// 	var taxRate float64

// 	fmt.Print("Revenue: ")
// 	fmt.Scan(&revenue)

// 	fmt.Print("Expences: ")
// 	fmt.Scan(&expences)

// 	fmt.Print("Tax Rate: ")
// 	fmt.Scan(&taxRate)

// 	ebt := revenue - expences
// 	profit := ebt * (1 - taxRate/100)
// 	ratio := ebt / profit

// 	fmt.Println(ebt)
// 	fmt.Println(profit)
// 	fmt.Println(ratio)
// }
