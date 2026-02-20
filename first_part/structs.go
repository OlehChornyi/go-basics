package main

// import (
// 	"fmt"

// 	"example.com/structs/user"
// )

// func main() {
// 	firstName := getUserData("Please enter your first name: ")
// 	lastName := getUserData("Please enter your last name: ")
// 	birthDate := getUserData("Please enter your birth date (MM/DD/YYYY): ")

// 	var appUser *user.User

// 	appUser, err := user.New(firstName, lastName, birthDate)

// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	appUser.PrintUserDetails()
// 	appUser.ClearUserName()
// 	appUser.PrintUserDetails()
// }

// func getUserData(promptText string) string {
// 	fmt.Print(promptText)
// 	var value string
// 	fmt.Scanln(&value)
// 	return value
// }
