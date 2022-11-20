package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func sumAtoB(a, b int) {
	sum := 0
	for ; a < b; a++ {
		sum += a
	}
	wg.Done()
	fmt.Println("SUM: ", sum)
}

func Gorutine() {
	wg.Add(4)
	for i := 0; i < 4; i++ {
		go sumAtoB(1, 100000)
	}
	wg.Wait()
	fmt.Println("끝입니다.")
}
