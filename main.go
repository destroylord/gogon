package main

import "fmt"

func main() {

	// refrensi : https://pkg.go.dev/fmt#pkg-overview
	name := "Dafrin"
	const tiket int = 50
	var remainingTickets uint = 50

	fmt.Printf("Hello World %v\n", name)
	fmt.Printf("Tiket untuk nonton bola: %v\n", tiket)

	var firstName string
	var lastname string
	var email string
	var userTiket uint

	fmt.Printf("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Printf("Enter your last name: ")
	fmt.Scan(&lastname)

	fmt.Printf("Enter your email: ")
	fmt.Scan(&email)

	fmt.Printf("Enter number of ticket: ")
	fmt.Scan(&userTiket)

	remainingTickets = remainingTickets - uint(userTiket)

	userTiket = 50
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation eamil at %v\n",
		firstName, lastname, userTiket, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, name)
}
