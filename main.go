package main

import (
	"fmt"
	"strings"
)

func main(){
	conferenceName := "Go Conference"
	const conferenceTickets  = 45
	var remainingTickets int = 45 
	var platformName = "Aryans"
	var firstName string
	var lastName string
	var userTickets int
	var bookings []string
	bookings = append(bookings, "STGI", "Barbosa")
	

	remainingTickets -= userTickets
	fmt.Printf(" *** Welcome to %v Booking Platform! ***\n", platformName)
	fmt.Println("Get your tickets here to attend " + conferenceName)
	fmt.Println("All tickets:", conferenceTickets, "Remaining:", remainingTickets, "\n")

	for i := 0; i < 3; i++ {
		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)
		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)

		
	}
	
	fmt.Printf("User %v %v booked %v tikcets.", firstName, lastName, userTickets)
}
