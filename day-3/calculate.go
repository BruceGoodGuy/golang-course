package main

import (
	"fmt"

	"example.com/day-3/fileops"
)

const fileName = "data.txt"

func checkBalance() {
	balance, _ := fileops.GetFloatFromFile(fileName)

	fmt.Printf("Your balance is %v\n", balance)
}

func depositMoney() {
	var amount float64 = 0
	fmt.Print("Enter deposit amount: ")
	fmt.Scan(&amount)
	if amount < 0 {
		amount = 0
	}
	balance, _ := fileops.GetFloatFromFile(fileName)
	balance += amount
	fileops.SaveFloatToFile(balance, fileName)
	fmt.Printf("Balance updated! New amount: %v\n", balance)
}

func withdrawMoney() {
	var amount float64 = 0
	balance, _ := fileops.GetFloatFromFile(fileName)
	fmt.Print("Withdrawal amount: ")
	fmt.Scan(&amount)

	if amount > balance {
		fmt.Println("Invalid amount. You can't withdraw more than you have.")
		return
	}
	balance -= amount
	fileops.SaveFloatToFile(balance, fileName)
	fmt.Printf("Balance updated! New amount: %v\n", balance)
}
