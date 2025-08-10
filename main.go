// Evertyhing in Go is structured in packages
// package - collection of a source files
package main

import "fmt"

// Main function is the entry point of our application
func main() {
	// We can use := to specify variable and assign it a value imediately
	// It doesn't work with constants, or when we want to specify a type explicitly
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Print("Please enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Please enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Please enter your email: ")
	fmt.Scan(&email)

	fmt.Print("Enter the number of tickets you want to purchase: ")
	fmt.Scan(&userTickets)

	fmt.Printf("Hello %v %v, thank you for booking %v tickets. You will receive a confirmation email at %v.", firstName, lastName, userTickets, email)
}
