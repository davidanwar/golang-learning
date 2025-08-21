package main

import "fmt"

func sayHello() {
	fmt.Println("Hello, World!")
}
func main() {
	sayHello()
	sayHello()
	sayHelloTo("Alice")
	hello := getHeloMessage("Bob")
	fmt.Println(hello)

	firstName, _ := getFullName()
	fmt.Println("First Name:", firstName)

}

func sayHelloTo(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func getHeloMessage(name string) string {
	return "Hello, " + name + "!"
}

func getFullName() (string, string) {
	return "John", "Doe"
}
