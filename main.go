package main

import (
	"fmt"
	"strings"
)

func main(){
	confrenceName := "Go Conference"
	const conferenceTickets uint8 = 50
	var remainingTickets uint8 = 50
	bookings := []string{}

	greetUsers(confrenceName, conferenceTickets, remainingTickets)



	for {
		// get user input
		firstName, lastName, email, userTickets := getUserInput()

		// validate user input
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(
			firstName,
			lastName,
			email,
			userTickets,
			remainingTickets,
		)

		if isValidName && isValidEmail && isValidTicketNumber {
			// update remaining tickets
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName + " " + lastName)

			fmt.Printf(
				"Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n",
				firstName, lastName, userTickets, email,
			)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, confrenceName)

			firstNames := getFirstNames(bookings)
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

func greetUsers(confName string, confTickets uint8, remTickets uint8) {
	fmt.Printf("Welcome to %v booking applciation\n", confName)
	fmt.Printf("we have a total of %v tickets, and %v are still available\n", confTickets, remTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings{
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(
	firstName string,
	lastName string,
	email string,
	userTickets uint8,
	remainingTickets uint8,
	) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets < remainingTickets
	
	return isValidName, isValidEmail, isValidTicketNumber
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