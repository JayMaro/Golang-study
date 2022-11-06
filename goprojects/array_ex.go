package main

import "fmt"

func main() {
	// 요소의 갯수 정의 필요 -> 정의하고 싶지 않다면 ...을 추가
	// 변수형 뒤에 []를 붙일 수 없음
	var intArray [5]int = [5]int{1, 2, 3, 4, 5}
	intArray2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(intArray)
	fmt.Println(intArray2)

	stringArray := [4]string{"mon", "tue"} // 값을 넣지 않으면 ""으로 초기화, int = 0, float = 0.0
	fmt.Println(stringArray)
	fmt.Println("stringArray[2] = ", stringArray[2], "stringArray[3] = ", stringArray[3])

	floatArray := [...]float64{1.1, 2.2, 3.3}
	fmt.Println(floatArray)

	//a := 5
	//floatArray2 := [a]float64{1.1, 2.2, 3.3} -> 배열의 크기는 항상 상수여야 한다.
	const b = 10
	floatArray3 := [b]float64{1.1, 2.2, 3.3}
	fmt.Println(floatArray3)

	stringArray[2] = "wed"
	for i := 0; i < len(stringArray); i++ {
		fmt.Println(stringArray[i])
	}

	for index, value := range stringArray {
		fmt.Println("index = ", index, "value = ", value)
	}

	x := [...]int{1, 2, 3, 4, 5}
	y := [...]int{10, 20, 30, 40, 50}
	y = x
	fmt.Println(x)
	fmt.Println(y)
	y[2] = 100
	// Golang의 배열의 대입은 값을 복사한다.
	fmt.Println(x) //[1 2 3 4 5]
	fmt.Println(y) //[1 2 100 4 5]

	//z := [5]int{1, 2, 3, 4, 5}
	//k := [7]int{10, 20, 30, 40, 50}
	//k = z -> 강타입 언어기 때문에 배열의 타입이 같더라도 크기가 다르면 연산되지 않는다.
	//y = [10]int{0, 0, 0, 0} -> 강타입 언어기 때문에 배열의 타입이 같더라도 크기가 다르면 연산되지 않는다.

	array2D := [2][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
	}
	for _, l := range array2D {
		for _, v := range l {
			fmt.Print(v, ", ")
		}
		fmt.Println()
	}
}
