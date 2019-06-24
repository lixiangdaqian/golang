package stri_test

import (
	"testing"
	"strings"
	"strconv"
)

func TestStringInit(t *testing.T) {
	str := "中国"
	t.Log(str)
	r := []rune(str)
	t.Logf("unicode:%x", r[0])
	t.Logf("utf8:%x", str)
}

func TestStringFunc(t *testing.T) {
	s := "a,b,c,d,e,f,g"
	repect := strings.Repeat(s, 2)
	t.Log("repect:"+repect)
	//按 , 隔开，返回切片
	parts := strings.Split(s, ",")
	t.Log(parts)
	join := strings.Join(parts, ":")
	t.Log(join)
}

func TestStringVert(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("string" + s)
	if in, err := strconv.Atoi(s); err == nil {
		t.Log(in + 10)
	}

}
