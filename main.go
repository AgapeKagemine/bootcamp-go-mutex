package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type Account struct {
	balance int64
	sync.Mutex
}

func addBalance(account *Account, wg *sync.WaitGroup, now time.Time) {
	defer wg.Done()
	account.Lock()
	account.balance += 100
	fmt.Println("ADD     : \t", account.balance, "\t", time.Since(now).Microseconds(), "ms")
	account.Unlock()
}

func deductBalance(account *Account, wg *sync.WaitGroup, now time.Time) {
	defer wg.Done()
	account.Lock()
	account.balance -= 100
	fmt.Println("DEDUCT  : \t", account.balance, "\t", time.Since(now).Microseconds(), "ms")
	account.Unlock()
}

func main() {
	now := time.Now()

	account := Account{
		balance: 1000,
		Mutex:   sync.Mutex{},
	}

	fmt.Println("INITIAL : \t", account.balance, "\t", time.Since(now).Microseconds(), "ms - with ", runtime.NumCPU(), " cpu")

	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(2)
		go addBalance(&account, &wg, now)
		go deductBalance(&account, &wg, now)
	}

	wg.Wait()

	fmt.Println("RESULT : \t", account.balance, "\t", time.Since(now).Microseconds(), "ms")
}
