// writting first code using Go

package main

import (
	"BOOKING-APP/helper"
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Conferance" // conferenceName := "Go Conferance" (alternative way to declare variable in go....but it will not work in constant)
const conferenceTickets = 50

var remainingTickets uint = 50
var bookings = make([]userData, 0)

type userData struct {
	firstName       string
	lastName        string
	email           string
	numberofTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	firstName, lastName, email, userTickets := getUserInput()

	isValidName, isValidEmail, isValidTicketNUmber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNUmber {

		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)
		//calling function for printing first name
		firstNames := printFirstNames()
		fmt.Printf("The first names of booking are: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Our conference is fully booked. Please come back next year.")

		}

	} else {

		if isValidName {
			fmt.Println("First name or last Name you enterted is too short")
		}

		if isValidEmail {
			fmt.Println("Your Email Address doesn't have @ sign")
		}

		if isValidTicketNUmber {
			fmt.Println("Number of ticket you entered is invalid")

		}

	}

	wg.Wait()
}

func greetUsers() {

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)

	fmt.Println("Get your tickets here to attend")

}

func printFirstNames() []string {

	firstNames := []string{}
	for _, booking := range bookings {

		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames

}

func getUserInput() (string, string, string, uint) {

	var firstName string
	var lastName string

	var email string

	var userTickets uint

	// ask user for their name (take input using fmt.scan) we also have to use & pointer before variable name in the function
	fmt.Println("Enter Your First Name")
	fmt.Scan(&firstName)

	fmt.Println("Enter Your last Name")
	fmt.Scan(&lastName)

	fmt.Println("Enter Your email")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets

}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {

	fmt.Printf("Thank You %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

	remainingTickets = remainingTickets - userTickets

	//create map for a user

	var userData = userData{

		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberofTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {

	time.Sleep(10 * time.Second)

	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")
	wg.Done()

}
