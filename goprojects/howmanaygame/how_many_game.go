package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100)
	var number int
	stdin := bufio.NewReader(os.Stdin)
	for true {
		if _, err := fmt.Scanln(&number); err != nil {
			stdin.ReadString('\n')
		}
		if target == number {
			fmt.Println("Correct")
			break
		} else if target > number {
			fmt.Println("Upper")
		} else {
			fmt.Println("Lower")
		}
	}

}
