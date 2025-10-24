package golangembed

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"os"
	"testing"
)

//go:embed version.txt
var version string

func TestString(t *testing.T) {
	fmt.Println(version)
}

//go:embed image.png
var image []byte

func TestSliceByte(t *testing.T) {
	err := os.WriteFile("image-new.png", image, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

//go:embed files/a.txt
//go:embed files/b.txt
var files embed.FS

func TestMultipleFiles(t *testing.T) {
	fileA, err := files.ReadFile("files/a.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(fileA))

	fileB, err := files.ReadFile("files/b.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(fileB))
}

//go:embed files/*.txt
var path embed.FS

func TestPathMatcher(t *testing.T) {
	entries, err := path.ReadDir("files")
	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			file, err := path.ReadFile("files/" + entry.Name())
			if err != nil {
				panic(err)
			}
			fmt.Println(entry.Name())
			fmt.Println(string(file))
		}
	}

}
