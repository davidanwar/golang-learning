package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blaclist Blacklist) {
	if blaclist(name) {
		fmt.Println("User is blacklisted:", name)
	} else {
		fmt.Println("User registered successfully:", name)
	}
}
func main() {
	blacklist := func(name string) bool {
		return name == "Anjing"
	}

	registerUser("Anjing", blacklist)
	registerUser("Eko", blacklist)
	registerUser("Budi", func(name string) bool {
		return name == "Anjing"
	})
}
