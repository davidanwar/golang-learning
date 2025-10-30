package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type People struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
	Married   bool   `json:"married"`
}

func TestJsonTag(t *testing.T) {
	customer := People{
		FirstName: "David",
		LastName:  "Anwar",
		Age:       31,
		Married:   true,
	}
	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
