package main

import "fmt"

func endApp() {
	fmt.Println("Ending the application")
	message := recover()
	if message != nil {
		fmt.Println("Error occurred:", message)
	}
}

func runApp(error bool) {
	defer endApp()
	fmt.Println("Running the application")
	if error {
		panic("Maaf Error")
	}
}
func main() {
	runApp(true)
	fmt.Println("Application finished running")
}
