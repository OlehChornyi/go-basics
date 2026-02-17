package main

import (
	"fmt"
)

// const inflationRate = 2.5

// func main() {
// 	var investmentAmount float64
// 	years, expectedReturnRate := 10.0, 5.5

// 	// fmt.Print("Investment Amount: ")
// 	outputText("Investment Amount: ")
// 	fmt.Scan(&investmentAmount)

// 	// fmt.Print("Expected Return Rate: ")
// 	outputText("Expected Return Rate: ")
// 	fmt.Scan(&expectedReturnRate)

// 	// fmt.Print("Investment Horizon(years): ")
// 	outputText("Investment Horizon(years): ")
// 	fmt.Scan(&years)

// 	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)
// 	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
// 	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

// 	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
// 	formattedRFV := fmt.Sprintf("Future Value (adjusted for infliation): %.2f\n", futureRealValue)
// 	// fmt.Println("Future Value:",futureValue)
// 	// fmt.Printf("Future Value: %.2f\nFuture Value (adjusted for infliation): %.2f", futureValue, futureRealValue)
// 	// fmt.Println("Future Value (adjusted for infliation):", futureRealValue)
// 	fmt.Print(formattedFV, formattedRFV)
// }

// func outputText(text string) { fmt.Print(text) }

// func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
// 	// fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
// 	// rfv := fv / math.Pow(1+inflationRate/100, years)
// 	// return fv, rfv
// 	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
// 	rfv = fv / math.Pow(1+inflationRate/100, years)
// 	return
// }

// Excercise 1: Profit Calculator
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

// Excercise 2: Profit Calculator refactoring
func main() {
	revenue := getUserInput("Revenue: ")
	expences := getUserInput("Expences: ")
	taxRate := getUserInput("Tax Rate: ")

	ebt, profit, ratio := calculateFinancials(revenue, expences, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.2f\n", profit)
	fmt.Printf("%.3f\n", ratio)
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

func calculateFinancials(revenue, expences, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expences
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return
}
