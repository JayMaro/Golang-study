package main

import (
	"fmt"
	"sync"
	"time"
)

type Car struct {
	Body  string
	Tire  string
	Color string
}

func MakeBody(bodyChan chan *Car, tireChan chan *Car) {
	for message := range bodyChan {
		time.Sleep(time.Second)
		message.Body = "BMW"
		tireChan <- message
	}
}

func MakeTire(tireChan chan *Car, colorChan chan *Car) {
	for message := range tireChan {
		time.Sleep(time.Second)
		message.Tire = "WITH CHAIN"
		colorChan <- message
	}
}

func PaintColor(colorChan chan *Car, wg sync.WaitGroup) {
	for message := range colorChan {
		if message.Body == "end" {
			wg.Done()
			return
		}
		time.Sleep(time.Second)
		message.Color = "BLUE"
		fmt.Printf("BODY: %s TIRE: %s COLOR: %s\n", message.Body, message.Tire, message.Color)
	}
}

// TODO 리팩토링 필요 deadLock 발생
func main() {
	var wg sync.WaitGroup
	bodyChannel := make(chan *Car)
	tireChannel := make(chan *Car)
	colorChannel := make(chan *Car)

	wg.Add(1)
	go MakeBody(bodyChannel, tireChannel)
	go MakeTire(tireChannel, colorChannel)
	go PaintColor(colorChannel, wg)
	for i := 0; i < 10; i++ {
		bodyChannel <- &Car{}
	}
	bodyChannel <- &Car{Body: "end"}
	wg.Wait()
}
