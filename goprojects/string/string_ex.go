package main

import (
	"fmt"
	"strings"
)

func main() {
	longString := `이 글은 매우 긴 글입니다.
띄어쓰기도 포함되어 있는 아주 무지막지한 글이죠 \t\n 도 포합되어 있습니당
"쌍 따옴표도 있고" '작은 따옴표도 있습니당.'`
	fmt.Println(longString)

	hello := "Hello 월드"
	runeHello := []rune(hello)
	fmt.Println(len(hello))     // 메모리 주소의 크기를 출력한다.
	fmt.Println(len(runeHello)) // 배열의 길이를 출력한다.

	for _, v := range hello {
		fmt.Printf("%c ", v)
	}
	fmt.Println()
	for i := 0; i < len(hello); i++ {
		fmt.Printf("%c ", hello[i]) // 바이트 단위로 출력해서 이상한 문자 출력됨
	}
	fmt.Println()
	for i := 0; i < len(runeHello); i++ {
		fmt.Printf("%c ", runeHello[i])
	}
	fmt.Println()
	// 문자열 비교 -> 문자열도 결국 숫자기 때문에 숫자의 크기로 비교한다.
	fmt.Println("ab" > "ac")

	// 문자열은 불변이기 때문에 새로운 문자열을 더해서 만들땐 새로운 주소에 문자열이 생성
	// -> 문자열을 계속 더하는건 쓸데없는 메모리 할당이 된다.
	// 해결하기 위해선 Builder를 사용
	I := "I"
	Like := "Like"
	Golang := "Golang"
	stringList := [...]string{I, Like, Golang}
	fmt.Println(I + " " + Like + " " + Golang)

	var stringBuilder strings.Builder

	for _, word := range stringList {
		for _, char := range word {
			stringBuilder.WriteRune(char)
		}
	}
	fmt.Println(stringBuilder.String())

	like := Like
	fmt.Println(like) // 문자열의 대입은 문자열 자체를 복사하는 것이 아닌 가르키는 대상을 동일하게 하는 것
}
