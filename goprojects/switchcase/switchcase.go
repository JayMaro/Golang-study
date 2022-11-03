package main

import "fmt"

func main() {
	a := 22

	switch a {
	case 1, 11:
		fmt.Println("1")
	case 2, 22:
		fmt.Println("2")
	default:
		fmt.Println("3")
	}

	b := 28

	switch { // true 생략 가능, if else 문처럼 조건으로 분기하여 쓸 수 있다.
	case b < 10:
		fmt.Println("Too Young")
	case b < 20:
		fmt.Println("Spring")
	case b < 30:
		fmt.Println("20s")
		fallthrough // 아래 case 까지 실행
	default:
		fmt.Println("Good Health is Very Important")
	}

	switch number := getNumber(); { // 초기문; 비굣값(아무것도 존재하지 않는다면 true)
	case number < 3:
		fmt.Println("low")
	case number == 3:
		fmt.Println("same")
	default:
		fmt.Println("high")
	}

}

func getNumber() int {
	return 3
}
