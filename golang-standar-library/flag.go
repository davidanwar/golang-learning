package main

import (
	"flag"
	"fmt"
)

func main() {
	var username = flag.String("username", "root", "Username for login")
	var password = flag.String("password", "password", "Password for login")
	var port = flag.Int("port", 8080, "Port to run the server on")
	flag.Parse()

	fmt.Println("Username:", *username)
	fmt.Println("Password:", *password)
	fmt.Println("Port:", *port)
}
