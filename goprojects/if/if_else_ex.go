package main

import "fmt"

func main() {
	light := "Red"

	if light == "Red" {
		fmt.Println("멈춘다.")
	} else if light == "Green" {
		fmt.Println("건넌다.")
	} else {
		fmt.Println("일단 멈춘다.")
	}

	spicy := "Hot"
	menu := "noodle"
	if menu == "ramen" {
		if spicy == "Hot" {
			fmt.Println("매운 라면")
		} else {
			fmt.Println("안매운 라면")
		}
	} else {
		if spicy == "Hot" {
			fmt.Println("매운 국수")
		} else {
			fmt.Println("안매운 국수")
		}
	}

	// if 초기문; 조건문 -> 독특한 형태니 알아두자!
	if hot, success := isHot(spicy); success && hot {
		fmt.Println("맵다!!!!")
	} else {
		fmt.Println("안맵다ㅎ")
	}

}

func isHot(spicy string) (bool, bool) {
	return spicy == "Hot", true
}
