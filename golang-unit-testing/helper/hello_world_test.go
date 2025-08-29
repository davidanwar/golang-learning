package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	fmt.Println("Before Unit Test")
	m.Run()
	fmt.Println("After Unit Test")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Alice")

	if result != "Hello Alice" {
		t.Error("Result is not 'Hello Alice', got:", result)
	}
}

func TestHelloAssert(t *testing.T) {
	result := HelloWorld("Alice")
	assert.Equal(t, "Hello Alice", result, "Result should be 'Hello Alice'")
}

func TestSubTest(t *testing.T) {
	t.Run("Alice", func(t *testing.T) {
		result := HelloWorld("Alice")
		assert.Equal(t, "Hello Alice", result, "Result should be 'Hello Alice'")
	})

	t.Run("Bob", func(t *testing.T) {
		result := HelloWorld("Bob")
		assert.Equal(t, "Hello Bob", result, "Result should be 'Hello Bob'")
	})
}

func TestTableHello(t *testing.T) {
	test := []struct {
		name     string
		request  string
		expected string
	}{
		{name: "Alice", request: "Alice", expected: "Hello Alice"},
		{name: "Bob", request: "Bob", expected: "Hello Bob"},
		{name: "Charlie", request: "Charlie", expected: "Hello Charlie"},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			result := HelloWorld(tt.request)
			assert.Equal(t, tt.expected, result, "Result should be "+tt.expected)
		})
	}
}
