package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const defaultBalance float64 = 1000

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("data.txt", []byte(balanceText), 0644)
}

func getBalance() (float64, error) {
	balanceBytes, err := os.ReadFile("data.txt")
	if err != nil {
		return defaultBalance, errors.New("can't read data file")
	}

	balanceStr := string(balanceBytes)
	balance, err := strconv.ParseFloat(balanceStr, 64)
	if err != nil {
		return defaultBalance, errors.New("can't parse string")
	}

	return balance, nil
}
