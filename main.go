package main

import "fmt"

func main() {
	// fmt.Println("Hello World")
	// fmt.Println("This is a test")
	// var anotherTest = "This is another test"
	// fmt.Println(anotherTest)

	const totalConferenceRoomTickets int = 100
	var remainingConferenceRoomTickets int = 50

	fmt.Printf("Conference room tickets %v\n", totalConferenceRoomTickets)
	fmt.Printf("Total: %v\n", totalConferenceRoomTickets)
	fmt.Printf("Remaining: %v\n", remainingConferenceRoomTickets)

	var userName string = "John Doe"
	var userTickets uint = 2
	fmt.Printf("User name %v booked %v tickets", userName, userTickets)
	// fmt.Printf("User tickets: %v\n", userTickets)
}
