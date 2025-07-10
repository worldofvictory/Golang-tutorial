package main

import (
	"fmt"
	//"log"
	"strings"
)

// qui scrivo variabili che saranno visibili a tutte le funzioni
var remainingTickets uint = 50 //uint only positive numbers
var bookings = []string{}

func main() {
	greetUsers(remainingTickets) //chiamo la funzione greetUsers, che stampa il numero di biglietti rimanenti
	//deve stare fuori  dal for, //altrimenti viene resettato ad ogni iterazione
	for {
		firstName, lastName, email, userTickets := getUserInput() //chiamata alla funzione getUserInput
		//passiamo il value di remainingTickets alla funzione greetUsers, perchè è stata segnat in entrata nella creazione dlla funzio
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)
		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of the bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("All tickets are sold out!")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short, please try again")
			}
			if !isValidEmail {
				fmt.Println("The email you entered is invalid, please try again")
			}
			if !isValidTicketNumber {
				fmt.Println("The number of tickets you entered is invalid, please try again")
			}
			fmt.Println("Your data input is invalid. Please try again")
		}
	}

}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}
func greetUsers(remainingTickets uint) { //passiamo qualsiasi nome come parametro
	fmt.Printf("Tickets remaining: %v\n", remainingTickets)
}
func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}
func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Printf("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Printf("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Printf("Enter your email: ")
	fmt.Scan(&email)

	fmt.Printf("How many tickets would you like to book? ")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}
func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)
	fmt.Printf("%v tickets remaining to Conference\n", remainingTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

}
