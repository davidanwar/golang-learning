package main

import "fmt"

func main() {
	name := "David"

	if name == "David" {
		fmt.Println("Hello David!")
	} else {
		fmt.Println("Hello Stranger!")
	}

	if length := len(name); length > 5 {
		fmt.Println("Your name is long.")
	} else {
		fmt.Println("Your name is short.")
	}
}
