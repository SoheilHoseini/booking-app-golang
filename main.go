package main

import "fmt"

func main(){
	conferenceName := "Go Conference"
	const conferenceTickets  = 45
	var remainingTickets int = 45
	var platformName = "Aryans"

	fmt.Printf(" *** Welcome to %v Booking Platform! ***\n", platformName)
	fmt.Println("Get your tickets here to attend " + conferenceName)
	fmt.Println("All tickets:", conferenceTickets, "Remaining:", remainingTickets)

	// print type of a variable using %T
	fmt.Printf("remainigTickets type is: %T\n", conferenceName)

	var userName string
	var userTickets int
	userName = "Shir Zard"

	fmt.Printf("User %v booked %v tikcets.", userName, userTickets)

}
