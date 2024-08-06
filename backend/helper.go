package main

import (
	"strings"
)

func userInputvalid(firstName string, lastName string, email string, userTickets int) (bool, bool, bool) {
	isValidusername := len(firstName) >= 2 && len(lastName) >= 2
	isValidemail := strings.Contains(email, "@")
	isValidtickets := userTickets > 0 && userTickets <= remainingTickets
	return isValidemail, isValidusername, isValidtickets

}
