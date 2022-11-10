package testpack

import "fmt"

var (
	a = c + b // c가 먼저 오지만 초기화가 먼저 되지 않는다.
	b = f()   // a 다음으로 초기화된다.
	c = f()   // c까지 모두 초기화가 된 이후에 a가 초기화 된다.
	d = 3
)

func init() {
	d++
	fmt.Println("init function", d)
}

// 전역변수가 모두 초기화 된 이후 실행된다.
func f() int {
	d++
	fmt.Println("f() d:", d)
	return d
}

// 대문자로 적힌 함수는 public
func PrintD() {
	fmt.Println("d: ", d)
}
