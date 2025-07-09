package main

import (
	"fmt"
	"strings"
)

func main() {
	var firstName string
	var lastName string
	var email string
	var remainingTickets int = 50
	var userTickets int
	bookings := []string{} //deve stare fuori  dal for, //altrimenti viene resettato ad ogni iterazione
	for {

		fmt.Printf("Tickets remaining: %v\n", remainingTickets)

		fmt.Printf("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Printf("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Printf("Enter your email: ")
		fmt.Scan(&email)

		fmt.Printf("How many tickets would you like to book? ")
		fmt.Scan(&userTickets)

		if userTickets > remainingTickets {
			fmt.Printf("Sorry %v, we only have %v tickets remaining.\n", firstName, remainingTickets)
			continue
		}

		fmt.Printf("User %v has booked %v tickets.\n", firstName, userTickets)
		fmt.Printf("Thank you %v %v for booking %v tickets. A confirmation email has been sent to %v.\n", firstName, lastName, userTickets, email)
		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)
		fmt.Printf("Tickets remaining: %v\n", remainingTickets)
		fmt.Printf("The array langs: %v\n", len(bookings))

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The first names of the bookings are: %v\n", firstNames)
		if remainingTickets == 0 {
			fmt.Println("All tickets are sold out!")
			break
		}
		//var bookings []string
		//var bookings [50]string
		//ask user for their name
		//bookings[0] = firstName + " " + lastName
		//fmt.Printf("The whole array: %v\n", bookings)
		//fmt.Printf("Type of array: %T\n", bookings)
		//fmt.Printf("The first value: %v\n", bookings[0])
	}

}
