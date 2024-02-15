// hands-on experience with Go's syntax, control structures, data types, slices, maps and structs

package main

import (
	"fmt"
	"strconv"
)

// TODO: Define a struct for FinanceRecord with fields for amount, description, category, and type (income or expense)
type FinanceRecord struct {
	// Add fields here
}

// TODO: Define global slice to hold finance records
var financeRecords []FinanceRecord

// TODO: Implement function to add a finance record
func addFinanceRecord(/* parameters */) {
	// Implement function
}

// TODO: Implement function to print summary (total income, total expenses, net balance)
func printSummary() {
	// Implement function
}

// TODO: Implement function to print records filtered by type (income/expense)
func printRecordsByType(/* parameters */) {
	// Implement function
}

// TODO: Implement function for basic financial calculations (e.g., estimate next month savings)
func calculateSavings(/* parameters */) {
	// Implement function
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
		case "2":
			// TODO: Implement viewing summary
		case "3":
			// TODO: Implement viewing records by type
		case "4":
			// TODO: Implement calculating savings
		case "5":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}
