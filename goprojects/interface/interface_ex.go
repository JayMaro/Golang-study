package main

import "fmt"

// -----------------------인터페이스 기분---------------------------
type Sample interface {
	String() string
	//String(int) string -> 이름이 겹치면 안됨
	//_(x int) -> 반드시 이름이 있어야 함
}

type Stringer interface { // 보통 메서드에 er을 붙여서 선언
	String() string
}

type Inter interface {
	Int() int
}

type Student struct {
	Name string
	No   int
}

func (s Student) String() string {
	return fmt.Sprintf("Hi my name is %s and I'm %d years old", s.Name, s.No)
}

func (s Student) Int() int {
	return s.No
}

type Actor struct {
	Name string
}

func (a Actor) String() string {
	return fmt.Sprintf("Hi my name is %s, Nice to meet you", a.Name)
}

// -----------------------인터페이스 심화---------------------------

type Reader interface {
	Read() (n int, err error)
	Close() error
}

type Writer interface {
	Write() (n int, err error)
	Close() error
}

// 중복되는 Close는 한 번만 정의된다.
// Reader의 기능과 Writer의 기능이 모두 포함되어 있어야 한다.
type ReadWriter interface {
	Reader
	Writer
}

// interface{} -> 모든 객체를 받을수 있다.
type allType interface {
}

func printVal(v allType) {
	switch t := v.(type) {
	case int:
		fmt.Println("정수형 입니다.")
	case float64:
		fmt.Println("실수형 입니다.")
	case string:
		fmt.Println("문자열 입니다.")
	default:
		fmt.Println("%T타입 입니다.", t)

	}
}

func main() {
	var stringer Stringer
	stringer = Student{"마로", 10}
	fmt.Println(stringer.String())
	printVal(10)
	printVal(10.1)
	printVal("test")
	printVal(stringer)

	// -----------------------인터페이스 타입 변환---------------------------
	student := stringer.(Student)
	fmt.Println(student.Name)
	//actor := stringer.(Actor)
	//fmt.Printf(actor.Name) // Error 발생 가르키는 인스턴스가 다르기 때문

	// 인터페이스끼리 타입 변환
	var inter Inter = stringer.(Inter)
	fmt.Println(inter.Int())
}
