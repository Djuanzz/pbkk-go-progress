package main

import "fmt"

func main() {
	fmt.Println("PBKK GOLANG SECTION")

	conferenceName := "PBKK"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Println("Welcome to %v bookin system", conferenceName)
	fmt.Println("We have total of %v tickets and %v are tickets available", conferenceTickets, remainingTickets)

}
