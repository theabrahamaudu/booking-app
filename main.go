package main

import (
	"fmt"
	"github.com/theabrahamaudu/booking-app/helper"
)

const conferenceTickets uint8 = 50
var conferenceName = "Go Conference"
var remainingTickets uint8 = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint8
}

func main(){

	greetUsers()

	for {
		// get user input
		firstName, lastName, email, userTickets := getUserInput()

		// validate user input
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(
			firstName,
			lastName,
			email,
			userTickets,
			remainingTickets,
		)

		if isValidName && isValidEmail && isValidTicketNumber {
			// update remaining tickets
			bookTicket(
				remainingTickets,
				userTickets,
				firstName,
				lastName,
				email,
			)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Come back next year")
				break
			}	
		} else {
			if !isValidName {
				fmt.Println("First name or last name is too short")
			}
			if !isValidEmail{
				fmt.Println("Email address does not contain @ sign")
			}
			if !isValidTicketNumber{
				fmt.Println("Number of tickets you entered is invalid")
			}
		}


	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking applciation\n", conferenceName)
	fmt.Printf("we have a total of %v tickets, and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings{
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}


func getUserInput() (string, string, string, uint8) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint8

	// ask user for their details
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket (
		remainingTickets uint8,
		userTickets uint8,
		firstName string,
		lastName string,
		email string,
	) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf(
		"Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n",
		firstName, lastName, userTickets, email,
	)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
