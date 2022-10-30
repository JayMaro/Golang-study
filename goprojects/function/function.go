package main

import "fmt"

func main() {
	a := 7
	var b int32 = 8
	c := Add(7, 8)
	fmt.Println(Add(7, 8))
	fmt.Println(Add32(7, 8)) // 정수형을 넣으면 자동으로 형변환을 해서 넣어주는 듯하다.
	fmt.Println(c)
	//fmt.Println(Add(a, b)) -> TypeError
	fmt.Println(Add(a, int(b)))
	// ------------------------
	d, isSuccess := multiReturnFunc(a, 3)
	fmt.Println(d, isSuccess)
	// isSuccess 이미 선언된 변수지만 := 이 사용가능한 이유는 f가 선언되지 않았기 때문
	f, isSuccess := multiVariableReturnFunc(a, 3)
	fmt.Println(f, isSuccess)
}

func Add(a int, b int) int {
	return a + b
}

func Add32(a int32, b int32) int32 {
	return a + b
}

func multiReturnFunc(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

func multiVariableReturnFunc(a, b int) (result int, isSuccess bool) {
	if b == 0 {
		result = 0
		isSuccess = false
		return
	}
	result = a / b
	isSuccess = true
	return
}
