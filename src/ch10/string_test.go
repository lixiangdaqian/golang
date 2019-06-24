package stri_test

import "testing"

func TestStringInit(t *testing.T) {
	str := "中国"
	t.Log(str)
	r := []rune(str)
	t.Logf("unicode:%x", r[0])
	t.Logf("utf8:%x", str)
}
