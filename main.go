package main

import (
	"fmt"
	"strings"

	"github.com/Djuanzz/pbkk-go-progress/helper"
)

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func main() {
	fmt.Println("PBKK GOLANG SECTION")

	conferenceName := "PBKK"
	const conferenceTickets = 10
	var remainingTickets uint = conferenceTickets
	bookings := []string{}

	fmt.Printf("Welcome to %v bookin system\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are tickets available\n", conferenceTickets, remainingTickets)

	// --- LOOPING
	for {
		firstName, lastName, email, userTickets := getUserInput()

		// --- VALIDATION INPUT
		validInput := helper.InputValidation(firstName, lastName, email, userTickets, remainingTickets)

		if validInput {
			remainingTickets -= userTickets

			fmt.Printf("Thank you %v %v for booking %v tickets for %v. You will receive a convirmation email at %v\n", firstName, lastName, userTickets, conferenceName, email)

			fmt.Printf("There are %v tickets remaining\n", remainingTickets)

			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Data type : %T\n", bookings)
			fmt.Printf("Data : %v\n", bookings)
			fmt.Printf("Data length : %v\n", len(bookings))
		} else {
			fmt.Println("There are many invalid inputs")
		}

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

	// --- SWITCH CASE
	fmt.Println("Tolong masukkan asal kota")
	var city string
	fmt.Scan(&city)

	city = strings.ToLower(city)

	switch city {
	case "jakarta", "depok", "bogor":
		fmt.Println("Halo Jabodetabek")
	case "solo", "semarang", "salatiga":
		fmt.Println("Halo Jawa Tengah")
	case "surabaya", "sidoharjo":
		fmt.Println("Halo Jawa Timur")
	case "bulan":
		fmt.Println("Halo Alien")
	default:
		fmt.Println("Maaf kamu dari mana yaa")
	}

}
