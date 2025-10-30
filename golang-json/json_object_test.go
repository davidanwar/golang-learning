package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName string
	LastName  string
	Age       int
	Married   bool
}

func TestJsonObject(t *testing.T) {
	customer := Customer{
		FirstName: "David",
		LastName:  "Anwar",
		Age:       31,
		Married:   true,
	}
	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
