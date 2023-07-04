package main

import "fmt"

func main() {
	var conf_name = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	fmt.Printf("Welcome to %v Booking application\n", conf_name)
	fmt.Printf("We have total tickets %v and %v are still remaining", conferenceTickets, conferenceTickets-remainingTickets)
	var userName string
	var userTickets int
	// ask user for their name

	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
}
