package main

import (
	"fmt"
	"togo/helper"
)

func main() {

	var conferenceName = "GO Conference"
	const noOfTickets = 50

	var remainingTickets uint = 50

	fmt.Println("Hello to ", conferenceName)
	fmt.Println("maxTickets:", noOfTickets)

	var username string
	fmt.Println("Memory address of username:", &username) // Pointer
	username = "Tom Hardy"
	validUsername := helper.Validate(username)
	fmt.Println("validUsername?:", validUsername)
	// fmt.Scan(&username)														// user input
	fmt.Println("Memory address of username:", &username) // Pointer
	fmt.Printf("User %v \n", username)
	fmt.Println("Remaining tickets", remainingTickets)

	var bookings [50]string // array 			with defined size
	bookings[0] = "ab"
	fmt.Println(bookings)

	var users []string                    // slice - Declare without size
	users = append(users, conferenceName) //			add elements to slice
	fmt.Println("Slice", users)

}

func greet() {
	fmt.Println("Welcome to first sub func")
}
