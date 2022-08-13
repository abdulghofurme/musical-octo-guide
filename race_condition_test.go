package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	sum := 0

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j < 100; j++ {
				sum++
			}
		}()
	}

	time.Sleep(10 * time.Second)
	fmt.Println("Sum =", sum)
}
