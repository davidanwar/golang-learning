package main

import "fmt"

func main() {
	person := map[string]string{}
	person["name"] = "Alice"
	person["age"] = "30"

	nameMap := map[string]string{
		"firstName": "John",
	}

	fmt.Println("Person Name: ", person["name"])
	fmt.Println("Name Map First Name: ", nameMap["firstName"])

	book := make(map[string]string)
	book["title"] = "Go Programming"
	book["author"] = "John Doe"
	fmt.Println("Book Title: ", book["title"])
	fmt.Println("Book Author: ", book["author"])

}
