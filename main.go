package main

import "fmt"

func main() {

	var firstName string
	var lastName string
	var email string
	var userTickets = 2
	//ask user for their name

	fmt.Printf("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Printf("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Printf("Enter your email: ")
	fmt.Scan(&email)
	fmt.Printf("User %v has booked %v tickets.\n", firstName, userTickets)
	fmt.Printf("Thank you %v %v for booking %v tickets. A confirmation email has been sent to %v.\n", firstName, lastName, userTickets, email)
}
