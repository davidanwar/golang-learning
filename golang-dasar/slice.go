package main

import "fmt"

func main() {
	names := []string{"Alice", "Bob", "Charlie", "David", "Eve", "Frank"}
	slices1 := names[2:4]
	slices2 := names[:]

	fmt.Println("Slice 1: ", slices1)
	fmt.Println("Slice 2: ", slices2)

	days := []string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	slice3 := days[5:7]
	slice3[0] = "Sabtu Baru"
	slice3[1] = "Minggu Baru"
	fmt.Println("Slice 3: ", slice3)
	slice4 := append(slice3, "Hari Baru")
	fmt.Println("Slice 4: ", slice4)
	fmt.Println("Days: ", days)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "New Day 1"
	newSlice[1] = "New Day 2"
	fmt.Println("New Slice: ", newSlice)
	fmt.Println("Length of New Slice: ", len(newSlice))
	fmt.Println("Capacity of New Slice: ", cap(newSlice))

	iniArray := [...]string{"A", "B", "C", "D", "E"}
	iniSlice := []string{"X", "Y", "Z"}
	fmt.Println("Initial Array: ", iniArray)
	fmt.Println("Initial Slice: ", iniSlice)

}
