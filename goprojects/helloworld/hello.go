// Go는 무조건 패키지가 존재
// package main -> 프로그램 시작점을 포함하는 패키지, main 함수가 존재
// main package는 한 프로그램당 하나가 존재, 둘 이상 선언할 수 없다.
package main

// 다른 패키지 import
import "fmt"

// 프로그램 시작점 ex(Java: public static void main)
func main() {
	// 불러온 패키지에서 함수 호출
	fmt.Println("Hello World")
}
