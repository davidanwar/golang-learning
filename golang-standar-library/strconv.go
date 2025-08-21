package main

import (
	"fmt"
	"strconv"
)

func main() {
	result, err := strconv.ParseBool("benar")
	if err != nil {
		fmt.Println("Error parsing boolean:", err.Error())
	} else {
		fmt.Println("Parsed boolean:", result)
	}
}
