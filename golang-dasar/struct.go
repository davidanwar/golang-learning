package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello() {
	fmt.Println("Hello", customer.Name)
}

func main() {
	var eko Customer
	fmt.Println(eko)
	eko.Name = "Eko"
	eko.Address = "Indonesia"
	eko.Age = 30

	fmt.Println(eko)

	joko := Customer{
		Name:    "Joko",
		Address: "Indonesia",
		Age:     25,
	}
	fmt.Println(joko)

	udin := Customer{"Udin", "Indonesia", 20}
	fmt.Println(udin)

	udin.sayHello()
}
