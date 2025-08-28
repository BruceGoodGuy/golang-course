package main

import "fmt"

func main() {

	fmt.Println("Profit calculator...")

	revenue, expenses, tax_rate := getUserInput()

	ebt, eft, ratio := calculateFinancials(revenue, expenses, tax_rate)

	fmt.Printf("Earning before tax (EBT): %v\n", ebt)

	fmt.Printf("Earnings After Tax (Profit): %v\n", eft)

	fmt.Printf("Ratio: %9.2f\n", ratio)

}

func calculateFinancials(revenue float64, expenses float64, tax_rate float64) (ebt float64, eft float64, ratio float64) {
	ebt = revenue - expenses
	eft = ebt * (1 - tax_rate)
	ratio = ebt / eft

	return ebt, eft, ratio
}

func getUserInput() (revenue float64, expenses float64, tax_rate float64) {

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax rate: ")
	fmt.Scan(&tax_rate)

	return revenue, expenses, tax_rate
}
