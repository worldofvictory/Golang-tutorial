package main

import "fmt"

func main() {
	var firstName string
	var lastName string
	var email string
	var remainingTickets int = 50
	var userTickets int
	for {
		bookings := []string{}

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
		bookings = append(bookings, firstName+" "+lastName)
		fmt.Printf("Tickets remaining: %v\n", remainingTickets)
		fmt.Printf("The array langs: %v\n", len(bookings))
		fmt.Printf("These are all the bookings: %v\n", bookings)

		//var bookings []string
		//var bookings [50]string
		//ask user for their name
		//bookings[0] = firstName + " " + lastName
		//fmt.Printf("The whole array: %v\n", bookings)
		//fmt.Printf("Type of array: %T\n", bookings)
		//fmt.Printf("The first value: %v\n", bookings[0])
	}
}
