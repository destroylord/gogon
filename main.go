package main

import "fmt"

func main() {

	// refrensi : https://pkg.go.dev/fmt#pkg-overview
	var name string = "Dafrin"
	const tiket int = 50

	fmt.Printf("Hello World %v\n", name)
	fmt.Printf("Tiket untuk nonton bola: %v\n", tiket)

	var userName string
	var userTiket int

	userName = "Tom"
	userTiket = 50
	fmt.Printf("User: %v booked %v.\n", userName, userTiket)
}
