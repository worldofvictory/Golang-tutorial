package main

import (
	"booking-app/helper"
	"fmt"
	"time"
)

const conferenceName = "Go Conference"
const conferenceTickets = 50

var remainingTickets uint = 50
var bookings = make([]UserDate, 0) //slice of maps//list of maps and size of list or initial size
type UserDate struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUsers()

	firstName, lastName, email, userTickets := getInputData() //non puo essere semplificato perchè arriva dal func main() che viene decitato dal utente
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		bookTicket(userTickets, firstName, lastName, email)

		go sendTicket(userTickets, firstName, lastName, email)
		firstNames := getFirstNames()
		fmt.Printf("There are all our bookings: %v\n", firstNames)
		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out. Come back next year!")
			//break
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
func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
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

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	//create a map for a user

	// create a UserDate struct for a user
	var userDate = UserDate{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userDate)
	fmt.Printf("List of bookings is %v\n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining to %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)                                                       // il tempo che prendiamo er spedire e generare per es pdf dal momento che abbiamo acquistato i biglietti
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName) //пишу так щоб те що зберігається в консолі я иогла взяти собі як рядок
	fmt.Println("##########################")
	fmt.Printf("Sending ticket %v to email address %v\n", ticket, email)
	fmt.Println("##########################")
}
