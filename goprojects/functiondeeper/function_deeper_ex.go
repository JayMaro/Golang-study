package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	// defer
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}
	// 역순으로 호출됨
	defer fmt.Println("반드시 호출됩니다.")
	defer f.Close()
	defer fmt.Println("파일을 닫았습니다.")

	fmt.Println("파일에 내용을 입력합니다..")
	fmt.Fprintln(f, "파일에 내용을 입력합니다.")

	// 함수 변수로 사용
	x := getOpFunc('*')(2, 4)
	fmt.Println(x)
	// 리터럴 함수는 외부의 값을 참조
	func(y int) {
		x += y
	}(100)
	fmt.Println(x)
}

func sum(numbers ...int) int { // 가변 인자 여러개
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}

type OpFunc func(int, int) int

// 리터럴 함수
func getOpFunc(char rune) OpFunc {
	if char == '*' {
		return func(a, b int) int {
			return a * b
		}
	} else if char == '-' {
		return func(a, b int) int {
			return a - b
		}
	} else {
		return func(a, b int) int {
			return a + b
		}
	}
}
