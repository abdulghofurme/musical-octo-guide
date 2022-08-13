package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Abdul Ghofur"
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(2 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Abdul Ghofur"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(2 * time.Second)
}

func OnlyInChannel(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Abdul Ghofur"
}

func OnlyOutChannel(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyInChannel(channel)

	OnlyOutChannel(channel)
	time.Sleep(2 * time.Second)
}

func BufferIn(channel chan<- string, number int) {
	channel <- fmt.Sprintf("Hello %v", number)
}

func BufferOut(channel <-chan string) {
	fmt.Println(<-channel)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 5)
	defer close(channel)

	for i := 0; i <= 100; i++ {
		go BufferIn(channel, i)
		go BufferOut(channel)
	}

	time.Sleep(5 * time.Second)
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string, 5)

	go func() {
		for i := 0; i <= 100; i++ {
			BufferIn(channel, i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}
}

func TestGoroutinesVSIteration(t *testing.T) {
	for i := 0; i <= 100; i++ {
		fmt.Println("Hello", i)
	}
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	defer close(channel1)
	channel2 := make(chan string)
	defer close(channel2)

	go BufferIn(channel1, 1)
	go BufferIn(channel2, 2)

	for i := 0; i < 2; {
		select {
		case data := <-channel1:
			fmt.Println("Data :", data)
			i++
		case data := <-channel2:
			fmt.Println("Data :", data)
			i++
		default:
			fmt.Println("Menunggu data")
		}
	}
}
