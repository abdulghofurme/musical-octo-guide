package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	time.Sleep(time.Millisecond * 500)
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Ups")

	time.Sleep(time.Second * 1)
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutines(t *testing.T) {
	for i := 0; i <= 100000; i++ {
		go DisplayNumber(i)
	}
	time.Sleep(5 * time.Second)
}
