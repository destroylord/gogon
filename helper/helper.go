package helper

import "strings"

func ValidateUserInput(firstName string, lastname string, email string, userTiket uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastname) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTiket > 0 && userTiket <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
