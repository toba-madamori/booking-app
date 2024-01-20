package main

import (
	"fmt"
	"strings"
)

func main() {
	var remainingTickets uint = 50
	conferenceName := "Golang Explosion"

	greetUsers(conferenceName)
	fmt.Println("Thank you for booking a ticket for the Go conference")

	fmt.Printf("At the moment %v tickets are still available\n", remainingTickets)

	for remainingTickets > 0 {
		var firstName string
		var lastName string
		var email string
		var userTickets int
		// ask user for their details
		fmt.Println("enter your first name")
		fmt.Scan(&firstName)

		fmt.Println("enter your last name")
		fmt.Scan(&lastName)

		fmt.Println("enter your email address")
		fmt.Scan(&email)

		fmt.Println("enter number of tickets: ")
		fmt.Scan(&userTickets)

		if userTickets > int(remainingTickets) {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v number of tickets\n", remainingTickets, userTickets)
			continue
		}
		remainingTickets -= uint(userTickets)

		bookings := []string{}
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Number of tickets remaining are: %v \n", remainingTickets)
		fmt.Printf("User %v booked %v tickets. \n", firstName, userTickets)

		firstNames := []string{}
		for _, booking := range bookings {
			names := strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("These are all the first names of our bookings: %v\n", firstNames)

		noTicketsRemaining := remainingTickets == 0
		if noTicketsRemaining {
			// End the program
			fmt.Println("Conference tickets have been sold out, come back next year.")
			break
		} else {
			fmt.Printf("Tickets: %v are still available\n", remainingTickets)
			continue
		}
	}

	city := "london"

	switch city {
	case "New York":
		fmt.Println("The city's name is New york")
	case "london":
		fmt.Println("The city's name is london")
	case "Berlin", "Mexico City":
		fmt.Println("The city's name is either Berlin or Mexico")
	default:
		fmt.Println("No valid city selected")
	}

}

func greetUsers(confName string) {
	fmt.Printf("Welcome to our %v conference\n", confName)
}
