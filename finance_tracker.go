// hands-on experience with Go's syntax, control structures, data types, slices, maps and structs

package main

import (
	"fmt"
)

// TODO: Define a struct for FinanceRecord with fields for amount, description, category, and type (income or expense)
type FinanceRecord struct {
	// Add fields here
	amount      int
	description string
	category    string
	rectype     string
}

// TODO: Define global slice to hold finance records
var financeRecords []FinanceRecord

// TODO: Implement function to add a finance record
func addFinanceRecord(amount int, description string, category string, rectype string) {
	// Implement function
	var financeRecord = new(FinanceRecord)
	financeRecord.amount = amount
	financeRecord.description = description
	financeRecord.category = category
	financeRecord.rectype = rectype
	financeRecords = append(financeRecords, *financeRecord)
}

// TODO: Implement function to print summary (total income, total expenses, net balance)
func printSummary() {
	// Implement function
	var income int
	var expenses int
	for i := 0; i < len(financeRecords); i++ {
		if financeRecords[i].rectype == "income" {
			income += financeRecords[i].amount
		} else {
			expenses += financeRecords[i].amount
		}
	}
	fmt.Println("income", income)
	fmt.Println("expenses", expenses)
	fmt.Println("net balance", income-expenses)
}

// TODO: Implement function to print records filtered by type (income/expense)
func printRecordsByType(rectype string) {
	// Implement function
	for i := 0; i < len(financeRecords); i++ {
		if financeRecords[i].rectype == rectype {
			fmt.Println(financeRecords[i].amount, financeRecords[i].description, financeRecords[i].category)
		}
	}
}

// TODO: Implement function for basic financial calculations (e.g., estimate next month savings)
func calculateSavings( /* parameters */ ) {
	// Implement function
	var income int
	var expenses int
	for i := 0; i < len(financeRecords); i++ {
		if financeRecords[i].rectype == "income" {
			income += financeRecords[i].amount
		} else {
			expenses += financeRecords[i].amount
		}
	}
	fmt.Println("Savings: ", income-expenses)
}

// TODO: Implement function to get user input and return as appropriate data type
func getUserInput(prompt string) string {
	fmt.Print(prompt)
	var input string
	fmt.Scanln(&input)
	return input
}

// TODO: Implement main logic for user interaction
func main() {
	for {
		fmt.Println("\nPersonal Finance Tracker")
		fmt.Println("1. Add Record")
		fmt.Println("2. View Summary")
		fmt.Println("3. View Records by Type")
		fmt.Println("4. Calculate Savings")
		fmt.Println("5. Exit")
		choice := getUserInput("Choose an option: ")

		switch choice {
		case "1":
			// TODO: Implement adding a record
			fmt.Println("enter details")
			var amount int
			var description string
			var category string
			var rectype string
			fmt.Scanln(&amount, &description, &category, &rectype)
			addFinanceRecord(amount, description, category, rectype)
		case "2":
			printSummary()
			// TODO: Implement viewing summary
		case "3":
			var rectype string
			fmt.Scanln(&rectype)
			printRecordsByType(rectype)
			// TODO: Implement viewing records by type
		case "4":
			calculateSavings()
			// TODO: Implement calculating savings
		case "5":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}
