package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func (u User) Output() {
	fmt.Println("Name: ", u.firstName+u.lastName, " DoB: ", u.birthDate)
}

func (u *User) Clear() {
	u.firstName = ""
	u.lastName = ""
	u.birthDate = ""
}

func (u User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func NewUser(firstName string, lastName string, birthDay string) (*User, error) {
	if firstName == "" || lastName == "" || birthDay == "" {
		return nil, errors.New("firstname, lastname and birthday must not empty")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDay,
	}, nil
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "Admin",
			lastName:  "Admin",
			birthDate: "----",
			createdAt: time.Now(),
		},
	}
}
