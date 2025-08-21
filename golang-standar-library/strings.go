package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Hello, World!", "World"))             // true
	fmt.Println(strings.ToUpper("hello"))                               // HELLO
	fmt.Println(strings.ToLower("HELLO"))                               // hello
	fmt.Println(strings.TrimSpace("  Hello, World!  "))                 // "Hello, World
	fmt.Println(strings.Trim("!!!Hello, World!!!", "!"))                // "Hello, World"
	fmt.Println(strings.Split("a,b,c", ","))                            // [a b c]
	fmt.Println(strings.Join([]string{"a", "b", "c"}, "-"))             // a
	fmt.Println(strings.ReplaceAll("Hello, World!", "World", "Gopher")) // Hello, Gopher!

}
