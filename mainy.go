package main

//import (
//	"fmt"
//	"strings"
//)

//func main() {
//	const conferenceName = "Go Conference"
//	const conferenceTickets = 50
//	var remainingTickets uint = 50
//var bookings []string

//	for {

//		var firstName string
//		var lastName string
//		var email string
//		var userTickets uint
//		var bookings = []string{}

//		fmt.Printf("Enter your first name: \n")
//		fmt.Scan(&firstName)

//		fmt.Printf("Enter your last name: \n")
//		fmt.Scan(&lastName)

//		fmt.Printf("Enter your email: \n")
//		fmt.Scan(&email)

//		fmt.Printf("How many tickets would you like to book? \n")
//		fmt.Scan(&userTickets)

//		isValidName := len(firstName) >= 2 && len(lastName) >= 2
//		isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
//		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

//		if isValidName && isValidEmail && isValidTicketNumber {
//			remainingTickets = remainingTickets - userTickets
//			bookings = append(bookings, firstName+" "+lastName)

//			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
//			fmt.Printf("%v tickets remaining to %v\n", remainingTickets, conferenceName)
//		firstNames := []string{}
//		for _, booking := range bookings {
//			var names = strings.Fields(booking)
//			firstNames = append(firstNames, names[0])
//		}
//		fmt.Printf("there are all our bookings: %v\n", firstNames)
//		if remainingTickets == 0 {
//			fmt.Println("Our conference is booked out. Come back next year!")
//			break
//		}
//	} else {
//		if !isValidName {
//			fmt.Println("First name or last name you entered is too short, please try again")
//		}
///		if !isValidEmail {
//			fmt.Println("The email you entered is invalid, please try again")
//		}
//		if !isValidTicketNumber {
//			fmt.Printf("We only have %v tickets remaining, you cannot book %v tickets\n", remainingTickets, userTickets)
//		}
//	}
//}
//}
//func greetUsers(confName string, confTickets uint, remainingTickets uint) {
//	fmt.Printf("Welcome to %v booking application\n", confName)
//	fmt.Printf("We have a total of %v tickets and %v are still available\n", confTickets, remainingTickets)

//}
