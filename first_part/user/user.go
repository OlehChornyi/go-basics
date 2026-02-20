package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName  string
	lastName   string
	birthDate  string
	creratedAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

// Data method
func (u User) PrintUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

// Mutation method
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

// Constructor - creation method - utility function
func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("First name, Last name & Birth date are required")
	}

	return &User{
		firstName:  firstName,
		lastName:   lastName,
		birthDate:  birthDate,
		creratedAt: time.Now(),
	}, nil
}

func NewAdmin(email, password string) *Admin {
	return &Admin{
		email:    email,
		password: password,
		User: User{
			firstName:  "firstName",
			lastName:   "lastName",
			birthDate:  "birthDate",
			creratedAt: time.Now(),
		}}
}
