package golanggoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{}
	pool.Put("Eko")
	pool.Put("Kurniawan")
	pool.Put("Khannedy")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)

		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Done")
}
