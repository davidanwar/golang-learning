package main

import "errors"

var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("not found errror")
)

func getById(id string) error {
	if id == "" {
		return ValidationError
	}

	// Simulating a not found scenario
	if id != "david" {
		return NotFoundError
	}

	return nil

}

func main() {
	err := getById("")

	if err != nil {
		if errors.Is(err, ValidationError) {
			println("Validation error occurred:", err.Error())
		} else if errors.Is(err, NotFoundError) {
			println("Not found error occurred:", err.Error())
		} else {
			println("An unexpected error occurred:", err.Error())
		}
	}

}
