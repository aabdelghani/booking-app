package main

import "fmt"

func main() {
	var conf_name = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	fmt.Printf("Welcome to %v Booking application\n", conf_name)
	fmt.Printf("We have total tickets %v and %v are still remaining", conferenceTickets, conferenceTickets-remainingTickets)
	var userName 
	// ask user for their name 

	userName = "Tom"
	fmt.println(userName)

}
