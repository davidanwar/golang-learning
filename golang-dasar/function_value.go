package main

import "fmt"

func getFunctionValue(value int) int {
	return value * value
}
func main() {

	kudarat := getFunctionValue
	fmt.Println("Nilai Kudarat:", kudarat(5))
}
