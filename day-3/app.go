package main

import (
	"fmt"

	"rsc.io/quote/v4"
)

func show() {
	var choice int
	var stop bool = false
	fmt.Print(`What do you want to do?
1. Check balance
2. Deposit money
3. Withdraw money
4. Exit
`)
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		checkBalance()
	case 2:
		depositMoney()
	case 3:
		withdrawMoney()
	case 4:
		stop = true
	default:
		stop = true
	}
	if stop {
		return
	}
	show()
}

func main() {
	fmt.Println(quote.Glass())
	show()
}
