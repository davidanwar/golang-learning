package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
func main() {

	total := sumAll(1, 2, 3, 4, 5)
	fmt.Println("Total:", total)

	numbers := []int{10, 20, 30, 40, 50}
	total = sumAll(numbers...)
	fmt.Println("Total with slice:", total)
}
