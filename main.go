package main

import (
	"fmt"
	"strings"
)

func main() {

	// refrensi : https://pkg.go.dev/fmt#pkg-overview
	name := "Dafrin"
	const tiket int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Hello World %v\n", name)
	fmt.Printf("Tiket untuk nonton bola: %v\n", tiket)

	for true{

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

		if userTiket <= remainingTickets {
			remainingTickets = remainingTickets - uint(userTiket)
			bookings = append(bookings, firstName+" "+lastname)

			fmt.Printf("Slice: %v\n", bookings)
			fmt.Printf("First Value: %v\n", bookings[0])
			fmt.Printf("Slice type: %T\n", bookings)
			fmt.Printf("Slice length: %v\n", len(bookings))

			userTiket = 50
			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation eamil at %v\n",
				firstName, lastname, userTiket, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, name)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				var firstName = names[0]
				firstNames = append(firstNames, firstName)
			}
			fmt.Printf("There first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {

			fmt.Printf("We only have %v ticket remaining, so you can't book %v tickets\n", remainingTickets, userTiket)
		}

	}
}
