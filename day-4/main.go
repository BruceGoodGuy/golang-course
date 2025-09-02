package main

import "fmt"

func main() {
	number := 1000
	var cloneNumber *int
	cloneNumber = &number

	// fmt.Println("clone address: ", cloneNumber)
	fmt.Println("Clone data: ", *cloneNumber)
	getAdultYears(&number)
	fmt.Println("New number use pointers: ", number)
}

func getAdultYears(age *int) {
	*age -= 18
}
