package helper

import "strings"

func InputValidation(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) bool {
	validName := len(firstName) >= 2 && len(lastName) >= 2
	validEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
	validTickets := userTickets > 0 && userTickets <= remainingTickets

	return validEmail && validName && validTickets
}
