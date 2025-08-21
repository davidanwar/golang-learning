package main

import "fmt"

func main() {
	var name [3]string

	name[0] = "Alice"
	name[1] = "Bob"
	name[2] = "Charlie"
	
	fmt.Println("Names in the array 0: ", name[0])
	fmt.Println("Names in the array 0: ", name[1])
	fmt.Println("Names in the array 0: ", name[2])

	var numbers = [3]int {
		10,
		20,
		30,
	}

	var values = [...]int {
		10,
		20,
		30,
		40,
		50,
	}

	fmt.Println("Numbers in the array numbers: ", numbers)
	fmt.Println("Panjang Array Number: ", len(numbers))
	fmt.Println("Panjang Array Values: ", len(values))
	
}