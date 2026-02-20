package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName  string
	lastName   string
	birthDate  string
	creratedAt time.Time
}

func (u user) printUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthDate := getUserData("Please enter your birth date (MM/DD/YYYY): ")

	var appUser user

	appUser = user{
		firstName:  firstName,
		lastName:   lastName,
		birthDate:  birthDate,
		creratedAt: time.Now(),
	}

	outputUserDetails(&appUser)
	appUser.printUserDetails()
}

func outputUserDetails(u *user) {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
