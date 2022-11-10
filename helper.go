package main

import "strings"

func validateUserInput(userFName string, userLName string, email string, userTickets uint) (bool, bool, bool){
	isValidName := len(userFName) >= 2 && len(userLName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNum := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNum
}