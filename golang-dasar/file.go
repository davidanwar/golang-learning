package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func createNewFile(name string, content string) error {
	file, error := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)

	if error != nil {
		return error
	}

	defer file.Close()
	file.WriteString(content)
	return nil
}

func readFile(name string) (string, error) {
	file, error := os.OpenFile(name, os.O_RDONLY, 0666)
	if error != nil {
		return "", error
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	var message string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		message += string(line) + "\n"
	}
	return message, nil
}

func main() {
	createNewFile("file.txt", "This is a new file created with Go.")
	result, _ := readFile("file.txt")
	fmt.Println(result)
}
