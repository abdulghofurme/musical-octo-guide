package golang_goroutines

import (
	"sync"
	"testing"
	"time"
)

func TestIteration(t *testing.T) {
	sum := 0

	for i := 1; i <= 1000; i++ {
		for j := 1; j <= 10; j++ {
			time.Sleep(1 * time.Millisecond)
			sum++
		}
	}

	t.Log("Sum =", sum)
}

func TestRaceConditionMutex(t *testing.T) {
	sum := 0
	var mutex sync.Mutex

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 10; j++ {
				time.Sleep(1 * time.Millisecond)
				mutex.Lock()
				sum++
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	t.Log("Sum =", sum)
}

func TestRaceCondition(t *testing.T) {
	sum := 0

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 10; j++ {
				time.Sleep(1 * time.Millisecond)
				sum++
			}
		}()
	}

	time.Sleep(5 * time.Second)
	t.Log("Sum =", sum)
}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(number int) {
	account.RWMutex.Lock()
	account.Balance += number
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() (balance int) {
	account.RWMutex.RLock()
	balance = account.Balance
	account.RWMutex.RUnlock()
	return
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				account.AddBalance(1)
				t.Log("Balance:", account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	t.Log("Total Balance:", account.GetBalance())
}
