package main

import "fmt"

func main() {
	counter := 0

	for counter < 10 {
		fmt.Println("Counter:", counter)
		counter++
	}

	for i := 0; i < 5; i++ {
		fmt.Println("Iteration:", i)
	}

	names := []string{"Alice", "Bob", "Charlie", "David", "Eve"}
	for index, name := range names {
		fmt.Printf("Index: %d, Name: %s\n", index, name)
	}
}
