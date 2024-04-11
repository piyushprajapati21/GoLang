package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferencetickets = 50
	var remainingTickets uint = 50

	greetUser(conferenceName, conferencetickets, remainingTickets)

	for {

		var firstName string
		var lastName string
		var email string
		var userTicket uint

		fmt.Println("\n\nEnter your First Name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your Last Name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your Email:")
		fmt.Scan(&email)

		fmt.Println("Enter number of Tickets:")
		fmt.Scan(&userTicket)

		isValidname := len(firstName) >= 2 && len(lastName) >= 2
		isEmail := strings.Contains(email, "@")
		isValidTicket := userTicket > 0

		if !isValidname || !isEmail || !isValidTicket {
			fmt.Println("Please enter valid info..")
			break
		}

		if userTicket > remainingTickets {
			fmt.Println("You have entered wrong tickets.")
			break
		}

		var bookings [50]string
		bookings[0] = firstName + " " + lastName

		remainingTickets = remainingTickets - userTicket

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive confirmation email at %v email id. Thanks....  \n", firstName, lastName, userTicket, email)

		fmt.Printf("Remaining Tickets : %v", remainingTickets)

		//fmt.Printf("The first name of booking are: %v\n", firstNames)
		if remainingTickets == 0 {
			fmt.Println("Our conference id Booked out...")
			break
		}

	}
}
func greetUser(conferenceName string, conferencetickets int, remainingTickets uint) {
	fmt.Println("Welcome to our conference.")

	fmt.Printf("Welcome to Our %v Booking Application.\n", conferenceName)
	fmt.Println("We have Totals", conferencetickets, " tickets and remaining tickets are", remainingTickets)
	fmt.Println("Get your Tickets here to attends..")
}
