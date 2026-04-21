package main

import "fmt"

func main() {
	var conferenceName string = "'Attendo Conference'"
	const conferenceTickets int = 50
	var usedTickets int = 0
	var remainingTickets int = conferenceTickets - usedTickets

	fmt.Println("welcome to", conferenceName, "booking application")
	fmt.Println("Get your tickets here to attend")
	fmt.Println("We have total of:", conferenceTickets, "tickets, and total tickets remaining are:", remainingTickets)
}
