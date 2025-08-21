package main

import "fmt"

func random() interface{} {
	return "OK"
}

func main() {
	result := random()
	switch value := result.(type) {
	case string:
		fmt.Println("Received a string:", value)
	case int:
		fmt.Println("Received an int:", value)
	default:
		fmt.Println("Received an unknown type")
	}
}
