package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)

	a := 10
	b := 20
	f := 327994389124.8123
	c := 123

	fmt.Print("a:", a, "b:", b)
	fmt.Println("a:", a, "b:", b)
	fmt.Printf("a: %d b: %d f: %f\n", a, b, f)
	fmt.Printf("a: %5d", a)
	fmt.Printf("a: %05d", b)
	fmt.Printf("a: %-5d\n", c)

	n, err := fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(n, err)
		stdin.ReadString('\n') // 개행 문자가 나올때까지 읽는다.
	} else {
		fmt.Println(n, a, b)
	}

	n, err = fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}

}
