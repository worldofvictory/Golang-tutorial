package main

import (
	"fmt"
)

func with() {
	const conferenceName = "Go Conference"
	const conferenceTickets = 50
	remainingTickets := 50
	//var bookings []string
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
}
