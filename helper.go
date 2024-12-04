package helper

import "strings"

func ValidateInput(firstName, lastName, email string, numTickets uint16, remainingTickets uint16) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := numTickets > 0 && numTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTickets
}
