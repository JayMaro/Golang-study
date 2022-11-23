package main

import (
	"fmt"
	"sync"
	"time"
)

func square2(wg *sync.WaitGroup, ch chan int) {
	for n := range ch {
		time.Sleep(time.Second)
		fmt.Printf("Square: %d\n", n*n)
	}
	wg.Done()
}

func Channel2() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square2(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}

	close(ch)
	wg.Wait()
}
