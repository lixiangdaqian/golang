package com_test

import "testing"

func TestCompute(t *testing.T) {
	a, b := 10, 12
	t.Logf("%d+%d=%d", a, b, a+b)
}

const (
	Readable = 1 << iota
	Writable
	Execuable
)

/*
数组比较需要维度相同才能比较，否则语法错误
*/
func TestArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	//b := [...]int{1,2,3,4,5}
	c := [...]int{1, 2, 3, 4}
	d := [...]int{5, 2, 3, 4}
	t.Log(a == d)
	t.Log(a == c)
	t.Log()
}

/*
按位置清零 &^  后面只要是1，对应位置就清零
*/
func TestConstTry1(t *testing.T) {
	a := 7 //0111
	a = a &^ Writable
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Execuable == Execuable)

}
