package main
import "fmt"

func main(){
	var conferenceName = "Go Conference"
	const conferenceTickets  = 45
	var remainingTickets = 45

	fmt.Println("Welcome to Aryans Booking Platform!")
	fmt.Println("Get your tickets here to attend " + conferenceName)
	fmt.Println("All tickets:", conferenceTickets, "Remaining:", remainingTickets)
}
