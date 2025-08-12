// Evertyhing in Go is structured in packages
// package - collection of a source files
package main

import (
	"fmt"
	"strings"
)

// Main function is the entry point of our application
func main() {
	// We can use := to specify variable and assign it a value imediately
	// It doesn't work with constants, or when we want to specify a type explicitly
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50 // uint stands for just positive whole numbers
	var bookings []string

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Print("Please enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Print("Please enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Print("Please enter your email: ")
		fmt.Scan(&email)

		fmt.Print("Enter the number of tickets you want to purchase: ")
		fmt.Scan(&userTickets)

		// Update the remaining tickets variable
		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Hello %v %v, thank you for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstnames := []string{}
		// _ (Blank identifier) => to ignore variable we don't want to use
		// 						=> in Go we need to make unused variables explicit
		for _, booking := range bookings {
			// strings.FIelds() => splits the string with white space as separator
			// 					=> returns a slice with the split elements
			var names = strings.Fields(booking)
			firstnames = append(firstnames, names[0])
		}

		fmt.Printf("The first name of bookings are: %v\n", firstnames)
	}

}
