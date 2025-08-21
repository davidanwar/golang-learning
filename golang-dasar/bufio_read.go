package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {
	input := strings.NewReader("This is a sample input string for testing\n with new lines and spaces\n")
	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		fmt.Println(string(line))
	}
}
