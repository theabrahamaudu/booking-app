package main

import (
	"fmt"
	"strings"
)

func main(){
	confrenceName := "Go Conference"
	const conferenceTickets int8 = 50
	var remainingTickets uint8 = 50
	bookings := []string{}

	fmt.Printf("ConferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, confrenceName)

	fmt.Printf("Welcome to %v booking applciation\n", confrenceName)
	fmt.Printf("we have a total of %v tickets, and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
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

		if userTickets < remainingTickets {
			// update remaining tickets
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName + " " + lastName)

			fmt.Printf(
				"Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n",
				firstName, lastName, userTickets, email,
			)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, confrenceName)

			firstNames := []string{}
			for _, booking := range bookings{
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Come back next year")
				break
			}	
		} else {
			fmt.Printf(
				"We only have %v tickets remaining, so you can't book %v tickets\n",
				remainingTickets, userTickets,
			)
		}


	}
}
