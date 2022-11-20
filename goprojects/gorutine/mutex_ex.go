package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

type Account struct {
	Balnace int
}

func DepositAndWithdraw(account *Account) {
	mutex.Lock()
	defer mutex.Unlock()

	if account.Balnace < 0 {
		panic(fmt.Errorf("잔액은 0보다 커야합니다"))
	}
	account.Balnace += 1000
	time.Sleep(time.Millisecond)
	account.Balnace -= 1000
}

func Mutex() {
	var wg sync.WaitGroup

	account := &Account{0}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				DepositAndWithdraw(account)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
