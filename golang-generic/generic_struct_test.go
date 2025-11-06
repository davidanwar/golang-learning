package golang_generic

import (
	"fmt"
	"testing"
)

type Data[T string] struct {
	Address T
	Name    T
}

func TestGenericStruct(t *testing.T) {
	data := Data[string]{
		Address: "Sukabumi",
		Name:    "David",
	}

	fmt.Println(data)
}
