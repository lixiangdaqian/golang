package try_test

import (
	"testing"
)



func TestFbnacci(t *testing.T) {
	var a, b = 1, 1
	t.Log(a)
	for i := 0; i < 10; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
}

func TestExchange(t *testing.T) {
	a, b := 1, 2
	t.Log(a, b)
	a, b = b, a
	t.Log(a, b)
}
