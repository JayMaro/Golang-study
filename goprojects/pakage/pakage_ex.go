package main

import (
	"fmt"
	"math/rand"
	"pakage/testpack"
	_ "pakage/testpkg"
)

func main() {
	fmt.Println(rand.Int())

	testpack.PrintD()

}
