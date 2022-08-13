package golang_goroutines

import (
	"testing"
	"time"
)

func RunHelloWorld(t *testing.T) {
	time.Sleep(time.Millisecond * 500)
	t.Log("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld(t)
	t.Log("Ups")

	time.Sleep(time.Second * 1)
}

func DisplayNumber(t *testing.T, number int) {
	t.Log("Display", number)
}

func TestManyGoroutines(t *testing.T) {
	for i := 0; i <= 100000; i++ {
		go DisplayNumber(t, i)
	}
	time.Sleep(5 * time.Second)
}
