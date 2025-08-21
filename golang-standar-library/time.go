package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("Current Time:", now)

	fmt.Println("Current Year:", now.Year())

	formattedTime := "2006-01-02 15:04:05"
	timeString := "2023-10-01 12:00:00"
	parsedTime, err := time.Parse(formattedTime, timeString)
	if err != nil {
		fmt.Println("Error parsing time:", err)
	} else {
		fmt.Println("Parsed Time:", parsedTime)
	}
}
