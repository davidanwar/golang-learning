package main

import "fmt"

type Filter func(string) string

func spamFilter(name string, filter Filter) {
	fmt.Println(filter(name))
}

func filter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return "Hello " + name
	}
}

func main() {
	spamFilter("Anjing", filter)

	filterDulu := filter
	spamFilter("Eko", filterDulu)
}
