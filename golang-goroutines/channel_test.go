package golanggoroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)
	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Eko"
		fmt.Println("Selesai mengirim data ke channel")
	}()
	data := <-channel
	fmt.Println(data)

	time.Sleep(2 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "David Anwar"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)
	go GiveMeResponse(channel)
	data := <-channel
	fmt.Println(data)
	time.Sleep(2 * time.Second)
}

func TestBufferChannel(t *testing.T) {
	channel := make(chan string, 2)
	defer close(channel)
	channel <- "Eko"
	channel <- "Kurniawan"

	fmt.Println("Mengirim data ke channel 1")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data ", data)
	}

	fmt.Println("Selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)
	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)
	couter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Menerima data dari channel 1 ", data)
			couter++
		case data := <-channel2:
			fmt.Println("Menerima data dari channel 2 ", data)
			couter++
		}
		if couter == 2 {
			break
		}
	}
}
