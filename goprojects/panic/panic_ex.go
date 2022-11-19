package main

import "fmt"

func divide(a, b int) {
	if b == 0 {
		panic("b는 0일수 없습니다.")
	}
	fmt.Printf("%d / %d = %d\n\n", a, b, a/b)
}

func recoverFunc() {
	fmt.Println("f() 함수 시작")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("패닉 복구", r)
		}
	}()

	divideFunc()
	fmt.Println("f() 함수 끝")
}

func divideFunc() {
	divide(9, 3)
	divide(9, 0)
}

func main() {
	recoverFunc()
	fmt.Println("recover 되어 계속 실행")
}
