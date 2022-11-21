package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg2 sync.WaitGroup

func Deadlock() {
	rand.Seed(time.Now().UnixNano())

	wg2.Add(2)
	fork := &sync.Mutex{}
	spoon := &sync.Mutex{}

	go diningProblem("A", fork, spoon, "포크", "수저")
	go diningProblem("B", spoon, fork, "수저", "포크")
	wg2.Wait()
}

func diningProblem(name string, first, second *sync.Mutex, firstName, secondName string) {
	for i := 0; i < 100; i++ {
		fmt.Printf("%s 밥을 먹으려 합니다.\n", name)
		first.Lock()
		fmt.Printf("%s 획득\n", firstName)
		second.Lock()
		fmt.Printf("%s 획득\n", secondName)
		fmt.Printf("%s 밥을 먹습니다.\n", name)

		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		second.Unlock()
		first.Unlock()
	}
	wg2.Done()
}
