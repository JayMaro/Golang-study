package main

import "fmt"

func main() {
	a := 3 // int - 64 bit = int64 -> 하지만 연산되지 않음
	var b float64 = 3.5
	var c int = int(b) // 3
	d := float64(a) * b

	var e int64 = 7
	f := a * int(e)

	fmt.Println(a, b, c, d, e, f)
}
