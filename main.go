package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	balance int64
	mut     *sync.Mutex
}

var now = time.Now()

func addBalance(account *Account, wg *sync.WaitGroup) {
	defer wg.Done()

	if account.mut.TryLock() {
		account.balance += 100
		fmt.Println("ADD    : \t", account.balance, "\t", time.Since(now).Microseconds(), "ms")
		account.mut.Unlock()
	} else {
		fmt.Println("ADD    :\t IS \t WAITING")
	}
}

func deductBalance(account *Account, wg *sync.WaitGroup) {
	defer wg.Done()

	if account.mut.TryLock() {
		account.balance -= 100
		fmt.Println("DEDUCT : \t", account.balance, "\t", time.Since(now).Microseconds(), "ms")
		account.mut.Unlock()
	} else {
		fmt.Println("DEDUCT :\t IS \t WAITING")
	}
}

func main() {
	var account = Account{
		balance: 1000,
		mut:     &sync.Mutex{},
	}

	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(2)
		go addBalance(&account, &wg)
		go deductBalance(&account, &wg)
	}

	wg.Wait()

}
