package main

import "fmt"

func main() {

	var firstName string
	var lastName string
	var email string
	var remainingTickets int = 50
	var userTickets int
	//ask user for their name
	fmt.Printf("Tickets remaining: %v\n", remainingTickets)

	fmt.Printf("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Printf("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Printf("How many tickets would you like to book? ")
	fmt.Scan(&userTickets)

	fmt.Printf("Enter your email: ")
	fmt.Scan(&email)

	fmt.Printf("User %v has booked %v tickets.\n", firstName, userTickets)
	fmt.Printf("Thank you %v %v for booking %v tickets. A confirmation email has been sent to %v.\n", firstName, lastName, userTickets, email)
	remainingTickets = remainingTickets - userTickets
	fmt.Printf("Tickets remaining: %v\n", remainingTickets)
}
