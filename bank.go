package main

// import (
// 	"fmt"

// 	"example.com/bank/fileops"
// 	"github.com/Pallinder/go-randomdata"
// )

// const accountBalanceFile = "balance.txt"
// const defaultBalanceValue = 1000

// func main() {
// 	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile, defaultBalanceValue)

// 	if err != nil {
// 		fmt.Println("ERROR")
// 		fmt.Println(err)
// 		fmt.Println("----------")
// 		panic("Can't continue. Sorry.")
// 	}

// 	fmt.Println("Welcome to Go Bank!")
// 	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())

// 	for {
// 		presentOprions()

// 		var choice int
// 		fmt.Print("Your choice: ")
// 		fmt.Scan(&choice)

// 		switch choice {
// 		case 1:
// 			fmt.Println("Your balance is", accountBalance)

// 		case 2:
// 			fmt.Print("Your deposit: ")
// 			var depositAmount float64
// 			fmt.Scan(&depositAmount)

// 			if depositAmount <= 0 {
// 				fmt.Println("Invalid ammount. Must be greater than 0.")
// 				continue
// 			}

// 			accountBalance += depositAmount
// 			fmt.Println("Balance updated! New amount:", accountBalance)
// 			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
// 		case 3:
// 			fmt.Print("Withdrawal amount: ")
// 			var withdrawAmount float64
// 			fmt.Scan(&withdrawAmount)

// 			if withdrawAmount <= 0 {
// 				fmt.Println("Invalid ammount. Must be greater than 0.")
// 				continue
// 			}

// 			if withdrawAmount > accountBalance {
// 				fmt.Println("Invalid ammount. You can't withdraw more than you have.")
// 				continue
// 			}

// 			accountBalance -= withdrawAmount
// 			fmt.Println("Balance updated! New amount:", accountBalance)
// 			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
// 		default:
// 			fmt.Println("Goodbye!")
// 			fmt.Println("Thanks for choosing our bank!")
// 			return
// 		}
// 	}

// }

// // Excercise 3: Profit Calculator refactoring
// // func main() {
// // 	revenue, err := getUserInput("Revenue: ")

// // 	if err != nil {
// // 		fmt.Println(err)
// // 		return
// // 	}

// // 	expences, err := getUserInput("Expences: ")

// // 	if err != nil {
// // 		fmt.Println(err)
// // 		return
// // 	}

// // 	taxRate, err := getUserInput("Tax Rate: ")

// // 	if err != nil {
// // 		fmt.Println(err)
// // 		return
// // 	}

// // 	ebt, profit, ratio := calculateFinancials(revenue, expences, taxRate)

// // 	fmt.Printf("%.1f\n", ebt)
// // 	fmt.Printf("%.2f\n", profit)
// // 	fmt.Printf("%.3f\n", ratio)
// // 	storeResults(ebt, profit, ratio)
// // }

// // func storeResults(ebt, profit, ratio float64) {
// // 	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRation: %.3f\n", ebt, profit, ratio)
// // 	os.WriteFile("results.txt", []byte(results), 0644)
// // }

// // func calculateFinancials(revenue, expences, taxRate float64) (ebt, profit, ratio float64) {
// // 	ebt = revenue - expences
// // 	profit = ebt * (1 - taxRate/100)
// // 	ratio = ebt / profit
// // 	return
// // }

// // func getUserInput(infoText string) (float64, error) {
// // 	var userInput float64
// // 	fmt.Print(infoText)
// // 	fmt.Scan(&userInput)

// // 	if userInput <= 0 {
// // 		return 0, errors.New("Value must be a positive number.")
// // 	}

// // 	return userInput, nil
// // }
