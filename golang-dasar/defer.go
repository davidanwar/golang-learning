package main

import "fmt"

func logging() {
	fmt.Println("Function execution finished")
}

func runExample() {
	fmt.Println("Running example function")
	defer logging()
}
func main() {
	runExample()
}
