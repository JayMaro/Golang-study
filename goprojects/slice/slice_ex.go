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

	array2 := [...]int{1, 2, 3, 4, 5, 6, 7}
	slicingSlice1 := array2[2:5] // 3, 4, 5 len: 3 cap: 5
	fmt.Println("slicingSlice1 = ", slicingSlice1, len(slicingSlice1), cap(slicingSlice1))
	slicingSlice2 := slicingSlice1[1:2] // 4 len: 1 cap: 4
	fmt.Println("slicingSlice2 = ", slicingSlice2, len(slicingSlice2), cap(slicingSlice2))
	slicingSlice3 := array2[2:5:5] // 3, 4, 5 len: 3 cap: 3
	fmt.Println("slicingSlice3 = ", slicingSlice3, len(slicingSlice3), cap(slicingSlice3))

	copySlice1 := []int{1, 2, 3, 4, 5, 6}
	copySlice2 := append([]int{}, copySlice1...)
	copySlice1[5] = 8
	fmt.Println("copySlice1 = ", copySlice1, len(copySlice1), cap(copySlice1))
	fmt.Println("copySlice2 = ", copySlice2, len(copySlice2), cap(copySlice2))

	copySlice3 := make([]int, 3, 10)
	copySlice4 := make([]int, 10, 10)
	cnt3 := copy(copySlice3, copySlice1)
	cnt4 := copy(copySlice4, copySlice1)
	fmt.Println("copySlice3 = ", copySlice3, len(copySlice3), cap(copySlice3))
	fmt.Println("cnt3 = ", cnt3)
	fmt.Println("copySlice4 = ", copySlice4, len(copySlice4), cap(copySlice4))
	fmt.Println("cnt4 = ", cnt4)

	// slice 중간요소 삭제
	removeSlice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	resultSlice1 := append(removeSlice1[:5], removeSlice1[6:]...)
	fmt.Println("resultSlice1 = ", resultSlice1, len(resultSlice1), cap(resultSlice1))

	// slice 중간요소 추가
	addSlice1 := append([]int{}, resultSlice1...)
	addSlice1 = append(addSlice1[:5], append([]int{6}, addSlice1[5:]...)...)
	fmt.Println("addSlice1 = ", addSlice1, len(addSlice1), cap(addSlice1))

}

func changeSliceNum(s []int) {
	s[1] = 500
}

func changeArrayNum(s [3]int) {
	s[1] = 500
}
