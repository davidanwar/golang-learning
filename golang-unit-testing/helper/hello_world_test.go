package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Alice with")

	if result != "Hello Alice" {
		t.Error("Result is not 'Hello Alice', got:", result)
	}
}
