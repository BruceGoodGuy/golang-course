package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	firstName := getUserData("Please enter your firstname: ")
	lastName := getUserData("Please enter your lastname: ")
	birthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User
	appUser, err := user.NewUser(firstName, lastName, birthDate)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}
	appUser.OutputUserDetails()
	appUser.Clear()
	appUser.Output()

	admin := user.NewAdmin("admin@admin.com", "1234")

	admin.OutputUserDetails()
	admin.Clear()
	admin.OutputUserDetails()
}

func getUserData(message string) string {
	fmt.Print(message)
	var data string
	fmt.Scan(&data)
	return data
}
