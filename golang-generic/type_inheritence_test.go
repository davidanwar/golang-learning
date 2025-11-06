package golang_generic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Employee interface {
	GetName() string
}

func GetName[T Employee](param T) string {
	name := param.GetName()
	fmt.Println(name)
	return name
}

type Manager interface {
	GetName() string
	GetManagerName() string
}

type MyManager struct {
	Name string
}

func (manager *MyManager) GetName() string {
	return manager.Name
}

func (manager *MyManager) GetManagerName() string {
	return manager.Name
}

func TestGetName(t *testing.T) {
	assert.Equal(t, "David", GetName[Manager](&MyManager{
		Name: "David",
	}))
}
