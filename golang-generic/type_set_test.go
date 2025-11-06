package golang_generic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Number interface {
	int | int32 | int64 | float32 | float64
}

func Min[T Number](first T, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func Max[T interface {
	int | int32 | int64 | float32 | float64
}](first T, second T) T {
	if first > second {
		return first
	} else {
		return second
	}
}

func TestMin(t *testing.T) {
	assert.Equal(t, 50, Min[int](100, 50))
	assert.Equal(t, int32(50), Min[int32](100, 50))
	assert.Equal(t, float64(50), Min[float64](100, 50))
}
