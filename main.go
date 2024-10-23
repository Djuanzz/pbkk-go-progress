package main

import "fmt"

func main() {
	fmt.Println("PBKK GOLANG SECTION")

	conferenceName := "PBKK"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to %v bookin system\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are tickets available\n", conferenceTickets, remainingTickets)

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Please enter your first name")
	fmt.Scanln(&firstName)

	fmt.Println("Please enter your last name")
	fmt.Scanln(&lastName)

	fmt.Println("Please enter your email")
	fmt.Scanln(&email)

	fmt.Println("Please enter the number of tickets you want to book")
	fmt.Scanln(&userTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets for %v. You will receive a convirmation email at %v\n", firstName, lastName, userTickets, conferenceName, email)

	remainingTickets -= userTickets
	fmt.Printf("There are %v tickets remaining\n", remainingTickets)

	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Data type : %T\n", bookings)
	fmt.Printf("Data : %v\n", bookings)
	fmt.Printf("Data length : %v\n", len(bookings))

}
