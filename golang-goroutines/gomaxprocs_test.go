package golanggoroutines

import (
	"runtime"
	"testing"
)

func TestGetGoMaxProcs(t *testing.T) {
	totalCpu := runtime.NumCPU()
	println("Total CPU ", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	println("Total Thread ", totalThread)

	totalGoroutines := runtime.NumGoroutine()
	println("Total Goroutines ", totalGoroutines)
}
