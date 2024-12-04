package main

import (
	"Booking-App/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets uint16 = 50

var conferenceName = "Go Conference"
var remainingTickets uint16 = 50
var bookings = make([]UserData, 0)

// var bookings = make([]map[string]string, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	greetUsers()

	if remainingTickets == 0 {
		fmt.Println("Our Conference is booked out. Please try again next year.")

	}

	firstName, lastName, email, numTickets := getUserInput()

	// Validation checks
	isValidName, isValidEmail, isValidTickets := helper.ValidateInput(firstName, lastName, email, numTickets, remainingTickets)
	if !isValidName {
		fmt.Println("First name and last name should each be at least 2 characters long.")

	}
	if !isValidEmail {
		fmt.Println("Please enter a valid email address.")

	}
	if !isValidTickets {
		fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets.\n", remainingTickets, numTickets)

	}

	// Booking is valid, proceed with booking
	bookTicket(firstName, lastName, email, numTickets)
	wg.Add(1)
	go sendTicket(uint(numTickets), firstName, lastName, email)
	displayFirstNames()
	wg.Wait()

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets, and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")
}

func displayFirstNames() {
	if len(bookings) > 0 {
		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)
	}
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint16) {
	var firstName, lastName, email string
	var numTickets uint16

	fmt.Print("Enter your First Name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your Last Name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your Email: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&numTickets)

	return firstName, lastName, email, numTickets
}

func bookTicket(firstName, lastName, email string, numTickets uint16) {
	remainingTickets -= numTickets

	//create a map for a user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: uint(numTickets),
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v for booking %v tickets.\nYou will receive a confirmation email at %v.\n", firstName, numTickets, email)
	fmt.Printf("Remaining tickets: %v\n", remainingTickets)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(3 * time.Second)
	var ticket = fmt.Sprintf("%v Ticket for %v %v", userTickets, firstName, lastName)
	fmt.Println("######################")
	fmt.Printf("Sending ticket: %v to email address: %v\n", ticket, email)
	fmt.Println("######################")

	wg.Done()
}
