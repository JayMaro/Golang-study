package main

import "fmt"

type User struct {
	name string
	no   int
}

func main() {
	value1 := 10
	value2 := 10
	var pointer1 *int = &value1
	pointer2 := &value1
	pointer3 := &value2

	fmt.Println(pointer1)
	fmt.Println(pointer2)
	fmt.Println(pointer3)
	fmt.Println(*pointer1)
	fmt.Println(*pointer2)
	fmt.Println(*pointer3)
	fmt.Println(*pointer1 == *pointer2)
	fmt.Println(*pointer1 == *pointer3)
	fmt.Println(pointer1 == pointer2)
	fmt.Println(pointer1 == pointer3)

	up1 := &User{"마시", 1}
	fmt.Println(up1)
	fmt.Println(*up1)
	changeValue(pointer1)
	fmt.Println(*pointer1)
	changeUser(up1)
	fmt.Println(*up1)
	up2 := new(User)
	fmt.Println(*up2)

	fmt.Println(returnAddress())

}

func changeValue(valuePointer *int) {
	*valuePointer = 30
}

func changeUser(userPointer *User) {
	userPointer.name = "마로"
	userPointer.no = 100
}

func returnAddress() *User {
	return new(User)
}
