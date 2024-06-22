package main

import (
	"fmt"
	"example.com/bank/fileops"
	"github.com/brianvoe/gofakeit/v7"
)

const accountBalanceFile = "balance.txt" 
var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

func main() {
	if err != nil {
		fmt.Println("Cannot find file")
	}

	fmt.Printf("Welcome to %v! \n", gofakeit.AppName()) 

	for {	
		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			showBalance()
		case 2:
			deposit()
			showBalance()
		case 3:
			withdraw()
			showBalance()
		default:
			fmt.Println("Bye!")
			return
		}
	}
}

func showBalance() {
	fmt.Println("Your balance is:", accountBalance)
}

func deposit() {
	var depositAmount float64
	fmt.Print("Your deposit amout: ")
	fmt.Scan(&depositAmount)
	accountBalance += depositAmount
	fileops.WriteFloatOnFile(accountBalanceFile, accountBalance)
}

func withdraw() {
	var withdrawAmount float64
	fmt.Print("Your withdraw amout: ")
	fmt.Scan(&withdrawAmount)
	accountBalance -= withdrawAmount
	fileops.WriteFloatOnFile(accountBalanceFile, accountBalance)
}
