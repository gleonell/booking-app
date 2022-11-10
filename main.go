package main

import "fmt"
//const mtTickets = 50
var meetupName = "Gophers meetup"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct{
	userFName string
	userLName string
	email string
	userTickets uint
}

func main() {
	
	greetUsers()

	//buying tickets
	for  {
		userFName, userLName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNum := validateUserInput(userFName, userLName, email, userTickets)

		if isValidEmail && isValidName && isValidTicketNum {
			bookTicket(userTickets, userFName, userLName, email)

			printFNames()
			
			if remainingTickets == 0 {
				fmt.Println("All tickets are sold")
				break
			}
		} else {
			if !isValidName {
				fmt.Printf("Invalid name data. Too short.\n")
			}
			if !isValidEmail {
				fmt.Printf("Invalid email address.\n")
			}
			if !isValidTicketNum {
				fmt.Printf("We have only %v tickets remaining. Or invalid num of tickets.\n", remainingTickets)
			}
		}
	}
	
}

func greetUsers() {
	fmt.Println("Welcome to my booking app for", meetupName)
	fmt.Println("Availible tickets: ", remainingTickets)

}

func printFNames() {
	firstNames := []string{}

	for _, booking := range bookings {
		
		firstNames = append(firstNames, booking.userFName)
	}

	fmt.Printf("This is all bookings list names: %v\n\n", firstNames)
}
//comtinue here


func getUserInput() (string, string, string, uint){
	var userFName string
	var userLName string
	var email string
	var userTickets uint


	fmt.Println("Enter your first name: ")
	fmt.Scan(&userFName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&userLName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return userFName, userLName, email, userTickets
}

func bookTicket(userTickets uint, userFName string, userLName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData {
		userFName: userFName,
		userLName: userLName,
		email: email,
		userTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", 
	userFName, userLName, userTickets, email)

	fmt.Printf("Remaining %v tickets.\n", remainingTickets)
}

