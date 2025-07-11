package main

import (
	"fmt"
	"strings"
)

func main() {
	const conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings = []string{}

	greetUsers(conferenceName, conferenceTickets, remainingTickets)
	for {
		firstName, lastName, email, userTickets := getInputData()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(remainingTickets, userTickets, bookings, firstName, lastName, email, conferenceName)
			firstNames := getFirstNames(bookings)
			fmt.Printf("There are all our bookings: %v\n", firstNames)
			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year!")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short, please try again")
			}
			if !isValidEmail {
				fmt.Println("The email you entered is invalid, please try again")
			}
			if !isValidTicketNumber {
				fmt.Printf("We only have %v tickets remaining, you cannot book %v tickets\n", remainingTickets, userTickets)
			}
		}
	}
}
func greetUsers(confName string, confTickets uint, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", confTickets, remainingTickets)
}

func getInputData() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Printf("Enter your first name: \n")
	fmt.Scan(&firstName)

	fmt.Printf("Enter your last name: \n")
	fmt.Scan(&lastName)

	fmt.Printf("Enter your email: \n")
	fmt.Scan(&email)

	fmt.Printf("How many tickets would you like to book? \n")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		names := strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidEmail, isValidName, isValidTicketNumber
}

func bookTicket(remainingTickets uint, userTickets uint, bookings []string, firstName string, lastName string, email string, conferenceName string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)
	fmt.Printf("List of bookings is %v\n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining to %v\n", remainingTickets, conferenceName)
}
