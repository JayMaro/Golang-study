package main

import "fmt"

// 코드값으로 사용
const Pig int = 0
const Cow int = 1
const Chicken int = 2

func main() {
	const C int = 10

	var b int = C * 20
	fmt.Println(b)
	// C = 20 -> 상수기 때문에 값을 저장할 수 없다.
	// fmt.Println(&C) -> 상수기 때문에 컴파일 타임에 리터럴로 전환되어 실행 파일에 값 형태로 쓰인다.
	// -> 메모리 영역을 사용하지 않음 -> 메모리 주소 존재 X

	PrintAnimal(Pig)
	PrintAnimal(Cow)
	PrintAnimal(Chicken)

	const (
		Red   int = iota
		Blue  int = iota
		Green int = iota
	)
	fmt.Println(Red, Blue, Green)

	const (
		C1 uint = iota + 1
		C2
		C3
	)
	fmt.Println(C1, C2, C3)

	const (
		BitFlag1 uint = 1 << iota
		BitFlag2
		BitFlag3
		BitFlag4
	)
	fmt.Println(BitFlag1, BitFlag2, BitFlag3, BitFlag4)

	// 타입 없는 상수
	const PI = 3.14
	var a int = 100 * PI
	fmt.Println(a)

}

func PrintAnimal(animal int) {
	if animal == Pig {
		fmt.Println("꿀꿀")
	} else if animal == Cow {
		fmt.Println("음머")
	} else if animal == Chicken {
		fmt.Println("꼬끼오")
	} else {
		fmt.Println("...?")
	}
}
