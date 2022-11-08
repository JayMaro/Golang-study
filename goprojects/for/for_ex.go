package main

import "fmt"

func main() {

	i := 0

	for i := 0; i < 10; i++ {
		fmt.Println(i, ", ")
	}

	for ; i < 10; i++ {
		fmt.Println(i, ", ")
	}

	for i := 0; i < 15; {
		fmt.Println("후처리 생략")
		fmt.Println(i, ", ")
		if i == 5 {
			break
		}
		i++
	}

	for i < 20 {
		fmt.Println("; 생략")
		fmt.Println(i, ", ")
		i++
	}

	for {
		fmt.Println("항상 true")
		i++
		if i < 25 {
			continue
		}
		fmt.Println(i, ", ")
		if i == 30 {
			break
		}
	}
}
