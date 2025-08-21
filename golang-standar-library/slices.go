package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Alice", "Bob", "Charlie", "David", "Eve"}
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(numbers))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(numbers))
	fmt.Println(slices.Contains(names, "Charlie")) // true
	fmt.Println(slices.Index(names, "David"))      // 3
}
