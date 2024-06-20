package main

import (
	"fmt"
	"sync"
)

var (
	balance int = 100
)

func Deposit(ammount int, wg *sync.WaitGroup, lock *sync.Mutex) {
	defer wg.Done()
	lock.Lock()
	b := balance
	balance = b + ammount
	lock.Unlock()
}

func Balance() int {
	b := balance
	return b
}

func main() {

	var wg sync.WaitGroup
	var lock sync.Mutex

	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go Deposit(i*100, &wg, &lock)
	}
	wg.Wait()
	fmt.Println(Balance())
}
