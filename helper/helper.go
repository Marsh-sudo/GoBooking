package helper
import "strings"

func ValidataUserInput(firstName string, lastName string, email string, userTickets uint,remainingTickets uint) (bool,bool,bool) {
	var isValidName = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail = strings.Contains(email, "@")
	var isValidTickets = userTickets > 0 && userTickets <= remainingTickets

	return isValidEmail,isValidName,isValidTickets

} 