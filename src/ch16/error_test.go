package error_test

import (
	"errors"
	"testing"
)

var LessError = errors.New("input must less than 20 and more than 1")

func Fibnacci(end int) ([]int, error) {
	if end >= 20 || end <= 1 {
		return nil, LessError
	}
	result := make([]int, 0, 10)
	var a, b = 1, 1
	result = append(result, b)
	for i := 1; i < end; i++ {
		tmp := a
		a = b
		b = tmp + a
		result = append(result, b)
	}
	return result, nil
}
func TestError(t *testing.T) {
	if res, err := Fibnacci(1); err == nil {
		t.Log(res)
	} else {
		t.Log("error: ", err)
	}
}
