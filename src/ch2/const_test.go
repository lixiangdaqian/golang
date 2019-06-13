package try_test

import "testing"

const (
	Monday    = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Sataday
	Sunday
)

const (
	Readable  = 1 << iota
	Writable
	Execuable
)

func TestConst(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday, Thursday, Friday, Sataday, Sunday)
}

func TestConstTry1(t *testing.T) {
	a := 8
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Execuable == Execuable)
}
