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

	// --- LOOPING
	for {
		fmt.Println("Please enter your first name")
		fmt.Scanln(&firstName)

		fmt.Println("Please enter your last name")
		fmt.Scanln(&lastName)

		fmt.Println("Please enter your email")
		fmt.Scanln(&email)

		fmt.Println("Please enter the number of tickets you want to book")
		fmt.Scanln(&userTickets)

		if userTickets > remainingTickets {
			fmt.Printf("Invalid input cuz there are only %v tickets\n", remainingTickets)
			continue
		}

		remainingTickets -= userTickets

		fmt.Printf("Thank you %v %v for booking %v tickets for %v. You will receive a convirmation email at %v\n", firstName, lastName, userTickets, conferenceName, email)

		fmt.Printf("There are %v tickets remaining\n", remainingTickets)

		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Data type : %T\n", bookings)
		fmt.Printf("Data : %v\n", bookings)
		fmt.Printf("Data length : %v\n", len(bookings))

		// --- CONDITIONAL
		if remainingTickets <= 0 {
			fmt.Println("All tickets are sold out")
			break
		}

	}

	// --- CONTOH LAIN DARI LOOPING
	kasus := 5
	for i := 0; i < kasus; i++ {
		fmt.Println("Kasus ke ", i)
	}

	books := []string{"Book 1", "Book 2", "Book 3"}
	for _, book := range books {
		fmt.Println(book)
	}

}
