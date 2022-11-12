package main

import "fmt"

func main() {
	array1 := [...]int{1, 2, 3}
	slice1 := []int{1, 2, 3}
	fmt.Printf("%T\n", array1)
	fmt.Printf("%T\n", slice1)

	changeSliceNum(slice1)
	changeArrayNum(array1)
	fmt.Println("slice1 = ", slice1, len(slice1), cap(slice1))
	fmt.Println("array1 = ", array1, len(array1), cap(slice1))

	slice2 := make([]int, 3)
	fmt.Println(slice2, len(slice2), cap(slice2))

	slice3 := append(slice2, 4)
	fmt.Println("slice2 = ", slice2, len(slice2), cap(slice2))
	// cap이 len과 같을 땐 cap을 2배로 늘리고 새로운 배열을 가르킨다.
	fmt.Println("slice3 = ", slice3, len(slice3), cap(slice3))

	slice4 := append(slice3, 5)
	fmt.Println("slice3 = ", slice3, len(slice3), cap(slice3))
	fmt.Println("slice4 = ", slice4, len(slice4), cap(slice4))

	// 슬라이스는 배열의 주소값을 가르키기 때문에 slice3과 slice4가 가르키는 배열이 같아서 이런 현상이 발생했다.
	slice3 = append(slice3, 100)
	fmt.Println("slice3 = ", slice3, len(slice3), cap(slice3))
	fmt.Println("slice4 = ", slice4, len(slice4), cap(slice4))

	// 여러 숫자들을 함께 추가할 수 있다.
	slice5 := append(slice4, 6, 7, 8, 9)
	fmt.Println("slice5 = ", slice5, len(slice5), cap(slice5))

}

func changeSliceNum(s []int) {
	s[1] = 500
}

func changeArrayNum(s [3]int) {
	s[1] = 500
}
