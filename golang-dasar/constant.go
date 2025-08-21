package main

import "fmt"

func main() {
	const name = "David Anwar"
	fmt.Println("Name: ", name)

	const (
		firstName = "John"
		lastName  = "Doe"
	)

	fmt.Println("Full Name: ", firstName, lastName)
}