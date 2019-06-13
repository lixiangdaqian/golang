package loop_test

import "testing"

func TestLoop(t *testing.T) {
	n := 0
	for n < 5 {
		n++
		t.Log(n)
	}

}

func TestIf(t *testing.T) {
	if n := 1 == 1; n {
		t.Log("haha")
	}
}

func TestSwitch(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("偶数")
		case i%2 == 1:
			t.Log("基数")
		default:
			t.Log("not 0-3")

		}
	}
}
