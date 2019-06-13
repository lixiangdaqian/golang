package typ_test

import (
	"testing"
	"math"
)

type MyInt int64

func TestDataType(t *testing.T) {
	var a int32 = 1
	var b int64
	b = int64(a)
	var c MyInt
	c = MyInt(b)
	t.Log(a, b, c)
}

func TestPreDefine(t *testing.T) {
	t.Log(math.MaxInt64)
}

/**
 指针不支持任何运算
 */
func TestPoint(t *testing.T) {
	a := 1
	point := &a
	t.Log(point)
	t.Logf("%T %T", a, point)
}

func TestString(t *testing.T) {
	var string string
	t.Log("*"+string+"*")
	t.Log(len(string))
}
