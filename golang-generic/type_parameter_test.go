package golang_generic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestSample(t *testing.T) {
	var result = Length[string]("david")
	fmt.Println(result)
	assert.Equal(t, "david", result)

	var resultNumber = Length[int](100)
	fmt.Println(result)
	assert.Equal(t, 100, resultNumber)

}
