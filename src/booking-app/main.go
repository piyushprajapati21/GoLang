package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferencetickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to Our %v Booking Application.\n", conferenceName)
	fmt.Println("We have Totals", conferencetickets, " tickets and remaining tickets are", remainingTickets)
	fmt.Println("Get your Tickets here to attends..")

	var firstName string
	var lastName string
	var email string
	var userTicket uint

	fmt.Println("Enter your First Name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your Last Name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your Email:")
	fmt.Scan(&email)

	fmt.Println("Enter number of Tickets:")
	fmt.Scan(&userTicket)

	remainingTickets = remainingTickets - userTicket

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive confirmation email at %v email id. Thanks....  \n", firstName, lastName, userTicket, email)

	fmt.Printf("Remaining Tickets : %v", remainingTickets)
}
