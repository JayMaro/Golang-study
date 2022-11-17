package main

import (
	"container/list"
	"fmt"
)

type Element struct {
	Value interface{}
	Next  *Element
	Prev  *Element
}

func main() {
	// 연결 리스트
	v := list.New()
	e4 := v.PushBack(4)
	e1 := v.PushFront(1)
	v.InsertBefore(3, e4)
	v.InsertAfter(2, e1)

	// 순회
	for e := v.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d ", e.Value)
	}

	for e := v.Back(); e != nil; e = e.Prev() {
		fmt.Printf("%d ", e.Value)
	}
	fmt.Println()

	m := make(map[string]int)
	m["첫째"] = 1
	m["둘째"] = 2
	m["셋째"] = 3
	m["넷째"] = 4
	m["다섯째"] = 5
	m["여섯째"] = 6

	delete(m, "여섯째")
	delete(m, "일곱째")
	for key, _ := range m {
		v, ok := m[key]
		fmt.Printf("값: %d 존재여부: %t \n", v, ok)
	}
	val, success := m["여섯째"]
	fmt.Printf("값: %d 존재여부: %t \n", val, success)
}
