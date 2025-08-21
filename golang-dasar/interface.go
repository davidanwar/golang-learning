package main

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

func sayHelloGuys(has HasName) {
	println("Hello ", has.GetName())
}

func (p Person) GetName() string {
	return p.Name
}
func main() {
	person := Person{Name: "Eko"}
	sayHelloGuys(person)
}
