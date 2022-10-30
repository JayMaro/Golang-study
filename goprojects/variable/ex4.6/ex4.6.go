package main

import "fmt"

var g int = 10 // 패키지 전역 변수

func main() {
	var m int = 20

	{
		var s int = 50
		fmt.Println(g, m, s)
	}

	printTest() // 같은 패키지 내의 메서드는 자유롭게 사용 가능!
}
