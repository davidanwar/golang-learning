package golang_generic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type GetterSetter[T any] interface {
	SetValue(value T) T
	GetValue() T
}

func ChangeValue[T any](param GetterSetter[T], value T) T {
	param.SetValue(value)
	return param.GetValue()
}

type MyData[T any] struct {
	Value T
}

func (data *MyData[T]) SetValue(value T) T {
	data.Value = value
	return data.Value
}

func (data *MyData[T]) GetValue() T {
	return data.Value
}

func TestInterface(t *testing.T) {
	data := &MyData[int]{}

	ChangeValue[int](data, 10)
	assert.Equal(t, 10, data.GetValue())
}
