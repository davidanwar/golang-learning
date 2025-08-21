package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type Users []User

func (u Users) Len() int {
	return len(u)
}

func (u Users) Less(i, j int) bool {
	return u[i].Age < u[j].Age
}

func (u Users) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func main() {
	users := []User{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
		{Name: "Fatih", Age: 10},
		{Name: "Anwar", Age: 15},
	}

	sort.Sort(Users(users))
	fmt.Println("Sorted users by age:")
	fmt.Println(users)
}
