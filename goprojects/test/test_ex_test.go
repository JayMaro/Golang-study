package main

import "testing"

func TestSquare(t *testing.T) {
	res := square(9)
	if res != 81 {
		t.Errorf("square(9) should be 81 but return %d", res)
	}
}

func TestSquare2(t *testing.T) {
	res := square(3)
	if res != 9 {
		t.Errorf("square(3) should be 9 but return %d", res)
	}
}
