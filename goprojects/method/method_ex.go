package main

import "fmt"

type Student struct {
	name  string
	no    int
	score Score
}

type Score struct {
	korean  int
	math    int
	english int
	science int
}

type Account struct {
	value int
}

type myInt int

func main() {
	student1 := Student{"마로", 1, Score{100, 100, 90, 100}}
	student2 := Student{"마시", 2, Score{90, 100, 70, 90}}
	fmt.Println(student1.score.getAvg())
	fmt.Println(student2.score.getAvg())

	a := myInt(10)
	fmt.Println(a.add(100))
	b := 20
	fmt.Println(myInt(b).add(100))

	acc1 := Account{100}
	acc2 := &Account{200}

	acc1.addValue(100) // 자동으로 변환
	acc2.addValue(200)
	fmt.Println(acc1)
	fmt.Println(acc2)
}

func (s Score) getAvg() int {
	return (s.korean + s.math + s.english + s.science) / 4
}

func (i myInt) add(num int) int {
	return num + int(i)
}

func (acc *Account) addValue(value int) {
	acc.value += value
}
